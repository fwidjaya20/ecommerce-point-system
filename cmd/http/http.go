package http

import (
	"github.com/fwidjaya20/ecommerce-point-system/cmd/container"
	userPointTransport "github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/transports/http"
	"github.com/fwidjaya20/ecommerce-point-system/lib/server"
	"github.com/go-chi/chi"
	kitHttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func MakeHandler(
	router *chi.Mux,
	container container.Container) http.Handler {
	opts := []kitHttp.ServerOption{
		kitHttp.ServerErrorEncoder(server.ErrorEncoder),
	}

	generateUserPointRoute(router, container, opts)

	return router
}

func generateUserPointRoute(router chi.Router, container container.Container, opts []kitHttp.ServerOption) {
	router.Group(func(r chi.Router) {
		r.Post("/point/{user_id}/add", userPointTransport.AddPoint(container.UserPointService, opts))
		r.Post("/point/{user_id}/deduct", userPointTransport.DeductPoint(container.UserPointService, opts))
		r.Get("/point/{user_id}/latest", userPointTransport.LatestPoint(container.UserPointService, opts))
	})
}