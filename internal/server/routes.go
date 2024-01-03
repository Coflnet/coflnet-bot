package server

import (
	"encoding/json"
	"fmt"
	"go.opentelemetry.io/otel"
	"log"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	tracer = otel.Tracer("api-tracer")
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", s.HelloWorldHandler)
	r.Post("/user/{uuid}", s.UpdateUserWithUUID)

	return r
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	_, span := tracer.Start(r.Context(), "hello-world-handler")
	defer span.End()

	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) UpdateUserWithUUID(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "update-user-with-uuid")
	defer span.End()

	uuid := chi.URLParam(r, "uuid")

	user, err := s.userService.LoadUserByUUID(ctx, uuid)
	if err != nil {
		span.RecordError(err)
		slog.Error("unable to load user", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	slog.Info(fmt.Sprintf("loaded user with external id %s", user.ExternalId))
	w.WriteHeader(http.StatusOK)
}
