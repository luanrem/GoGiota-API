package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (api *Api) BindRoutes() {
	api.Router.Use(middleware.RequestID, middleware.RealIP, middleware.Logger, middleware.Recoverer)

	api.Router.Route("/conta", func(r chi.Router) {
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("post created")
		})
	})
}
