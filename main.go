package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alexcuse/todo-authn/api"
	"github.com/alexcuse/todo-authn/db"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/keratin/authn-go/authn"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
)

type contextKey string

const userIDContextKey contextKey = "user"

var ErrNotAuthenticated = errors.New("user not authenticated")

func main() {
	logger := zerolog.New(os.Stdout)

	port := os.Getenv("TODO_PORT")
	addr := net.JoinHostPort("0.0.0.0", port)
	authHost := os.Getenv("AUTHN_URL")
	authPrivateHost := os.Getenv("AUTHN_PRIVATE_URL")
	dbHost := os.Getenv("DATABASE_URL")

	flag.Parse()

	dbPool, err := pgxpool.New(context.Background(), dbHost)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Use(cors.AllowAll().Handler)

	// setup handler for swagger.JSON
	swagger, err := api.GetSwagger()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to get swagger")
	}
	// Skip validating Server names
	swagger.Servers = nil
	swaggerJSON, err := swagger.MarshalJSON()
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to marshal swagger")
	}
	r.Handle("/swagger.json", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write(swaggerJSON)
	}))

	// setup API handler and auth middleware
	internalHandler := api.NewServer(db.New(dbPool), &logger, func(ctx context.Context) (int32, error) {
		if id := ctx.Value(userIDContextKey); id != nil {
			return id.(int32), nil
		}
		return 0, ErrNotAuthenticated
	})
	authClient, err := authn.NewClient(authn.Config{
		Issuer:         authHost,
		PrivateBaseURL: authPrivateHost,
		Audience:       "*",
		Username:       "authn",
		Password:       "authn",
		KeychainTTL:    120,
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to initialize authn client")
	}
	oapiHandler := api.NewStrictHandler(internalHandler, nil)
	apiHandler := api.Handler(oapiHandler)
	apiHandler = authMiddleware(authClient, &logger)(apiHandler)

	r.Mount("/api/v1/", apiHandler)

	s := &http.Server{
		Handler:           r,
		Addr:              addr,
		ReadHeaderTimeout: 10 * time.Second,
	}
	logger.Info().Str("address", addr).Msg("starting server")
	log.Fatal(s.ListenAndServe())
}

// authMiddleware provides a middleware function to check headers
// and cookies for a JWT token and store the associated user ID in
// context.
func authMiddleware(client *authn.Client, log *zerolog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			// only authenticate API requests
			if strings.HasPrefix(r.URL.Path, "/api/") {
				token := extractToken(r)
				if token == "" {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				subject, err := client.SubjectFrom(token)
				if err != nil {
					log.Error().Err(err).Msg("failed to verify claims")
					w.WriteHeader(http.StatusUnauthorized)
					return

				}

				userID, err := strconv.Atoi(subject)
				if err != nil {
					log.Error().Err(err).Msg("failed to parse user ID")
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				ctx = context.WithValue(r.Context(), userIDContextKey, int32(userID))
			}
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// extractToken will get the token from a given request.
// It will look first in the Authorization header and then in
// a cookie named authn.
func extractToken(r *http.Request) string {
	// Get the "Authorization" header value.
	rawToken := r.Header.Get("Authorization")

	if rawToken == "" {
		c, e := r.Cookie("todo")
		if e != nil && !errors.Is(e, http.ErrNoCookie) {
			panic(e)
		}
		if c != nil {
			rawToken = c.Value
		}
	}
	// Check if the header is not empty and starts with "Bearer ".
	if rawToken != "" && strings.HasPrefix(rawToken, "Bearer ") {
		return strings.TrimPrefix(rawToken, "Bearer ")
	}

	return rawToken
}
