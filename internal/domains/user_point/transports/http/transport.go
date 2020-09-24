package http

import (
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/endpoints"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/models"
	"github.com/fwidjaya20/ecommerce-point-system/lib/server"
	kitHttp "github.com/go-kit/kit/transport/http"
	"net/http"
)

func AddPoint(service user_point.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.Add(service), server.HTTPOption{
			DecodeModel: &models.PublishAddOrDeductPointRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func DeductPoint(service user_point.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.Deduct(service), server.HTTPOption{
			DecodeModel: &models.PublishAddOrDeductPointRequest{},
		}, opts).ServeHTTP(w, r)
	}
}

func LatestPoint(service user_point.UseCase, opts []kitHttp.ServerOption) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server.NewHTTPServer(endpoints.LatestPoint(service), server.HTTPOption{
			DecodeModel: &models.GetPointRequest{},
		}, opts).ServeHTTP(w, r)
	}
}
