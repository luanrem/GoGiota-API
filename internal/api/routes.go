package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer)

	api.Router.Route("/conta", func(r chi.Router) {
		r.Post("/", api.handleCriarConta)
	})
}
