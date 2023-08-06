package route

import (
	"net/http"

	handlers "dmoniac/src/interface/rest/handlers/riwayat"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func DmoniacRouter(h handlers.RiwayatHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/", h.Create)
	r.Get("/", h.GetList)

	return r
}
