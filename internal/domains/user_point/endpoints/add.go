package endpoints

import (
	"context"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/models"
	"github.com/fwidjaya20/ecommerce-point-system/lib/transport/http"
	"github.com/go-kit/kit/endpoint"
)

func Add(service user_point.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, error error) {
		payload := request.(*models.PublishAddOrDeductPointRequest)

		response, error = service.PublishAddPoint(ctx, *payload)

		return http.Response(ctx, response, nil), error
	}
}