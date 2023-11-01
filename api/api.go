package api

import (
	"context"

	"github.com/alexcuse/todo-authn/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog"
)

var _ StrictServerInterface = &Server{}

// NewServer creates a new server.
func NewServer(queries *db.Queries, log *zerolog.Logger, extractUserID func(context.Context) (int32, error)) Server {
	return Server{
		queries:       queries,
		log:           log,
		extractUserID: extractUserID,
	}
}

// Server defines user implementation of the generated HTTP server.
type Server struct {
	queries       *db.Queries
	log           *zerolog.Logger
	extractUserID func(context.Context) (int32, error)
}

func apiTodo(item db.Todo) Todo {
	var (
		id          *string
		title       *string
		description *string
		completed   *bool
	)

	id = &item.ID
	title = &item.Title
	completed = &item.Completed

	if item.Description.Valid {
		description = &item.Description.String
	}

	return Todo{
		Id:          id,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

// GetTodos returns the user's TODOs.
func (s Server) GetTodos(ctx context.Context, _ GetTodosRequestObject) (GetTodosResponseObject, error) {
	userID, err := s.extractUserID(ctx)
	if err != nil {
		return nil, err
	}
	todos, err := s.queries.GetUserTODOs(ctx, userID)
	if err != nil {
		s.log.Error().Err(err).Msg("failed to get user todos")
		return nil, err
	}

	var resp GetTodos200JSONResponse
	for _, t := range todos {
		resp = append(resp, apiTodo(t))
	}
	return resp, nil
}

// PostTodos creates a new TODO.
func (s Server) PostTodos(ctx context.Context, request PostTodosRequestObject) (PostTodosResponseObject, error) {
	userID, err := s.extractUserID(ctx)
	if err != nil {
		return nil, err
	}

	var title string
	if request.Body.Title != nil {
		title = *request.Body.Title
	}
	todo, err := s.queries.CreateTODO(ctx, db.CreateTODOParams{
		UserID: userID,
		Title:  title,
	})

	if err != nil {
		s.log.Error().Err(err).Msg("failed to create todo")
		return nil, err
	}

	return PostTodos201JSONResponse(apiTodo(todo)), nil
}

// DeleteTodosId deletes a TODO by Id.
func (s Server) DeleteTodosId(ctx context.Context, request DeleteTodosIdRequestObject) (DeleteTodosIdResponseObject, error) {
	userID, err := s.extractUserID(ctx)
	if err != nil {
		return nil, err
	}
	err = s.queries.DeleteTODO(ctx, db.DeleteTODOParams{
		ID:     request.Id,
		UserID: userID,
	})

	if err != nil {
		s.log.Error().Err(err).Msg("failed to delete todo")
		return nil, err
	}

	return DeleteTodosId204Response{}, nil
}

// GetTodosId returns a TODO by Id.
func (s Server) GetTodosId(ctx context.Context, request GetTodosIdRequestObject) (GetTodosIdResponseObject, error) {
	userID, err := s.extractUserID(ctx)
	if err != nil {
		return nil, err
	}
	todo, err := s.queries.GetTODO(ctx, db.GetTODOParams{
		ID:     request.Id,
		UserID: userID,
	})
	if err != nil {
		s.log.Error().Err(err).Msg("failed to get user todo")
		return nil, err
	}
	return GetTodosId200JSONResponse(apiTodo(todo)), nil
}

// PutTodosId updates a TODO by Id.
func (s Server) PutTodosId(ctx context.Context, request PutTodosIdRequestObject) (PutTodosIdResponseObject, error) {
	userID, err := s.extractUserID(ctx)
	if err != nil {
		return nil, err
	}

	var description pgtype.Text
	if request.Body.Description != nil {
		description = pgtype.Text{
			String: *request.Body.Description,
			Valid:  true,
		}
	}

	todo, err := s.queries.UpdateTODO(ctx, db.UpdateTODOParams{
		ID:          request.Id,
		UserID:      userID,
		Title:       request.Body.Title,
		Description: description,
		Completed:   request.Body.Completed,
	})

	if err != nil {
		s.log.Error().Err(err).Msg("failed to update todo")
		return nil, err
	}

	return PutTodosId200JSONResponse(apiTodo(todo)), nil
}
