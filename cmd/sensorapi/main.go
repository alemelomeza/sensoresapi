package main

import (
	"net/http"

	"github.com/alemelomeza/sensor-api/internal/repositories/subscribersrepo"

	"github.com/alemelomeza/sensor-api/internal/core/services/measuressrv"
	"github.com/alemelomeza/sensor-api/internal/core/services/subscriberssrv"
	"github.com/alemelomeza/sensor-api/internal/handlers/measureshdl"
	"github.com/alemelomeza/sensor-api/internal/handlers/subscribershdl"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	measuresService := measuressrv.NewService()
	measuresHandler := measureshdl.NewHTTPHandler(measuresService)
	subscribersRepository := subscribersrepo.NewRepository()
	subscribersService := subscriberssrv.NewService(subscribersRepository)
	subscribersHandler := subscribershdl.NewHTTPHandler(subscribersService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/v1", func(r chi.Router) {
		r.Post("/measures", measuresHandler.Create)
		r.Post("/subscribers", subscribersHandler.Create)
	})

	http.ListenAndServe(":80", r)
}
