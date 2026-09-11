package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/luanrem/GoGiota-API/internal/services"
)

type Api struct {
	Router       *chi.Mux
	ContaService services.ContaService
}
