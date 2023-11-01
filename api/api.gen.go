// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.2 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Todo defines model for Todo.
type Todo struct {
	Completed   *bool   `json:"completed,omitempty"`
	Description *string `json:"description,omitempty"`
	Id          *string `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// PostTodosJSONBody defines parameters for PostTodos.
type PostTodosJSONBody struct {
	Title *string `json:"title,omitempty"`
}

// PutTodosIdJSONBody defines parameters for PutTodosId.
type PutTodosIdJSONBody struct {
	Completed   bool    `json:"completed"`
	Description *string `json:"description,omitempty"`
	Title       string  `json:"title"`
}

// PostTodosJSONRequestBody defines body for PostTodos for application/json ContentType.
type PostTodosJSONRequestBody PostTodosJSONBody

// PutTodosIdJSONRequestBody defines body for PutTodosId for application/json ContentType.
type PutTodosIdJSONRequestBody PutTodosIdJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get All TODOs
	// (GET /todos)
	GetTodos(w http.ResponseWriter, r *http.Request)
	// Create a TODO
	// (POST /todos)
	PostTodos(w http.ResponseWriter, r *http.Request)
	// Delete a TODO by ID
	// (DELETE /todos/{id})
	DeleteTodosId(w http.ResponseWriter, r *http.Request, id string)
	// Get a TODO by ID
	// (GET /todos/{id})
	GetTodosId(w http.ResponseWriter, r *http.Request, id string)
	// Update a TODO by ID
	// (PUT /todos/{id})
	PutTodosId(w http.ResponseWriter, r *http.Request, id string)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get All TODOs
// (GET /todos)
func (_ Unimplemented) GetTodos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Create a TODO
// (POST /todos)
func (_ Unimplemented) PostTodos(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Delete a TODO by ID
// (DELETE /todos/{id})
func (_ Unimplemented) DeleteTodosId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Get a TODO by ID
// (GET /todos/{id})
func (_ Unimplemented) GetTodosId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update a TODO by ID
// (PUT /todos/{id})
func (_ Unimplemented) PutTodosId(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetTodos operation middleware
func (siw *ServerInterfaceWrapper) GetTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTodos(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostTodos operation middleware
func (siw *ServerInterfaceWrapper) PostTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostTodos(w, r)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// DeleteTodosId operation middleware
func (siw *ServerInterfaceWrapper) DeleteTodosId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.DeleteTodosId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// GetTodosId operation middleware
func (siw *ServerInterfaceWrapper) GetTodosId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetTodosId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PutTodosId operation middleware
func (siw *ServerInterfaceWrapper) PutTodosId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "id" -------------
	var id string

	err = runtime.BindStyledParameterWithLocation("simple", false, "id", runtime.ParamLocationPath, chi.URLParam(r, "id"), &id)
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "id", Err: err})
		return
	}

	ctx = context.WithValue(ctx, BearerAuthScopes, []string{})

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PutTodosId(w, r, id)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/todos", wrapper.GetTodos)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/todos", wrapper.PostTodos)
	})
	r.Group(func(r chi.Router) {
		r.Delete(options.BaseURL+"/todos/{id}", wrapper.DeleteTodosId)
	})
	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/todos/{id}", wrapper.GetTodosId)
	})
	r.Group(func(r chi.Router) {
		r.Put(options.BaseURL+"/todos/{id}", wrapper.PutTodosId)
	})

	return r
}

type GetTodosRequestObject struct {
}

type GetTodosResponseObject interface {
	VisitGetTodosResponse(w http.ResponseWriter) error
}

type GetTodos200JSONResponse []Todo

func (response GetTodos200JSONResponse) VisitGetTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostTodosRequestObject struct {
	Body *PostTodosJSONRequestBody
}

type PostTodosResponseObject interface {
	VisitPostTodosResponse(w http.ResponseWriter) error
}

type PostTodos201JSONResponse Todo

func (response PostTodos201JSONResponse) VisitPostTodosResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	return json.NewEncoder(w).Encode(response)
}

type DeleteTodosIdRequestObject struct {
	Id string `json:"id"`
}

type DeleteTodosIdResponseObject interface {
	VisitDeleteTodosIdResponse(w http.ResponseWriter) error
}

type DeleteTodosId204Response struct {
}

func (response DeleteTodosId204Response) VisitDeleteTodosIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(204)
	return nil
}

type DeleteTodosId404Response struct {
}

func (response DeleteTodosId404Response) VisitDeleteTodosIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type GetTodosIdRequestObject struct {
	Id string `json:"id"`
}

type GetTodosIdResponseObject interface {
	VisitGetTodosIdResponse(w http.ResponseWriter) error
}

type GetTodosId200JSONResponse Todo

func (response GetTodosId200JSONResponse) VisitGetTodosIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetTodosId404Response struct {
}

func (response GetTodosId404Response) VisitGetTodosIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

type PutTodosIdRequestObject struct {
	Id   string `json:"id"`
	Body *PutTodosIdJSONRequestBody
}

type PutTodosIdResponseObject interface {
	VisitPutTodosIdResponse(w http.ResponseWriter) error
}

type PutTodosId200JSONResponse Todo

func (response PutTodosId200JSONResponse) VisitPutTodosIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PutTodosId404Response struct {
}

func (response PutTodosId404Response) VisitPutTodosIdResponse(w http.ResponseWriter) error {
	w.WriteHeader(404)
	return nil
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get All TODOs
	// (GET /todos)
	GetTodos(ctx context.Context, request GetTodosRequestObject) (GetTodosResponseObject, error)
	// Create a TODO
	// (POST /todos)
	PostTodos(ctx context.Context, request PostTodosRequestObject) (PostTodosResponseObject, error)
	// Delete a TODO by ID
	// (DELETE /todos/{id})
	DeleteTodosId(ctx context.Context, request DeleteTodosIdRequestObject) (DeleteTodosIdResponseObject, error)
	// Get a TODO by ID
	// (GET /todos/{id})
	GetTodosId(ctx context.Context, request GetTodosIdRequestObject) (GetTodosIdResponseObject, error)
	// Update a TODO by ID
	// (PUT /todos/{id})
	PutTodosId(ctx context.Context, request PutTodosIdRequestObject) (PutTodosIdResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHttpHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHttpMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetTodos operation middleware
func (sh *strictHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	var request GetTodosRequestObject

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetTodos(ctx, request.(GetTodosRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTodos")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetTodosResponseObject); ok {
		if err := validResponse.VisitGetTodosResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostTodos operation middleware
func (sh *strictHandler) PostTodos(w http.ResponseWriter, r *http.Request) {
	var request PostTodosRequestObject

	var body PostTodosJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostTodos(ctx, request.(PostTodosRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostTodos")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostTodosResponseObject); ok {
		if err := validResponse.VisitPostTodosResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// DeleteTodosId operation middleware
func (sh *strictHandler) DeleteTodosId(w http.ResponseWriter, r *http.Request, id string) {
	var request DeleteTodosIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.DeleteTodosId(ctx, request.(DeleteTodosIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "DeleteTodosId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(DeleteTodosIdResponseObject); ok {
		if err := validResponse.VisitDeleteTodosIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// GetTodosId operation middleware
func (sh *strictHandler) GetTodosId(w http.ResponseWriter, r *http.Request, id string) {
	var request GetTodosIdRequestObject

	request.Id = id

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetTodosId(ctx, request.(GetTodosIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetTodosId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetTodosIdResponseObject); ok {
		if err := validResponse.VisitGetTodosIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PutTodosId operation middleware
func (sh *strictHandler) PutTodosId(w http.ResponseWriter, r *http.Request, id string) {
	var request PutTodosIdRequestObject

	request.Id = id

	var body PutTodosIdJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PutTodosId(ctx, request.(PutTodosIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PutTodosId")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PutTodosIdResponseObject); ok {
		if err := validResponse.VisitPutTodosIdResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/7xVT28aPxD9Ktb8fsdVIG1OeyNFreglqKHqIeJgdgdwtGu79myqVcR3r2ZMCMvSBtom",
	"J7xjz5/33szwCIWrvbNoKUL+CLFYY63lOHOl418fnMdABsXKzyskLPmDWo+Qw8K5CrWFTQYlxiIYT8bZ",
	"vQeRgrErvjflUTMZqvDIDV8li1vcY0FiiVg0wVB7y7Wmqq5RBwyjhtb8tZCvjy7UmiCHz99mkCVkUq3c",
	"wi7ymsjDhgMbuxTEHRAwUtEwZjWaTtTSBVVrq1fGrtTsZnyjDGEdL2CHAcQ6mk4ggwcMMQW5vBheDBmp",
	"82i1N5DDezFl4DWtBcSAXOnktELq1/EFKRh8QKVVZSIpt1S6qg6KYKk0O0xKyOET0kxiZhAwemdjouvd",
	"cJi0tIRWUmnvK1OI5+A+JvFSL/BJwvPh/4BLyOG/wXPXDLYtM5B+eRZMh6DbxGsXx21TFBjjsqnUU1Ed",
	"WSG/6wp6N9/MM4hNXevQJlRqtEUeOaN38QhfHwJqYrYs/nhmqU/S1MU9lr43GOnale1ZBHVn5DfdzAlM",
	"4Om542fbfBXCvN/pnecUGtz0ZLw8q8qX1eurJcQVQmWp4k66qj1Ps50YHE9cU7cPHk25SdrxVumrOBb7",
	"1lEUVItWGYpqMu5rmV6LmpOy3/VX/fgSNiU/xJfB1S89rCO1dI0tz6Ohi2bRqsmY87w07ydhfxr2Y8CH",
	"r94nR6f6FRjk4T+kz+ugayQMUdwN5+CtChlYLTvflHA4TNke4MM5nWfgmyOSfPWlPr0Zp01HkH+xWf7y",
	"3/f0xVQhZHvZ/mw7Dd9mOzWiyhtMb1f+bfu9GGHzMwAA//+pzqJHZAkAAA==",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}