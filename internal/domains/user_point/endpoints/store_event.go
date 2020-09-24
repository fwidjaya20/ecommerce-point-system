package endpoints

import (
	"context"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/models"
	"github.com/fwidjaya20/ecommerce-point-system/internal/globals"
	"github.com/fwidjaya20/ecommerce-point-system/lib/database"
	"github.com/fwidjaya20/ecommerce-point-system/lib/transport/http"
	"github.com/go-kit/kit/endpoint"
)

func StoreEvent(service user_point.UseCase) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, _ error) {
		payload := request.(*models.AddOrDeductPointRequest)

		errTrans := database.RunInTransaction(ctx, globals.DB(), func(ctx context.Context) error {
			return service.StoreEvent(ctx, *payload)
		})

		return http.Response(ctx, response, nil), errTrans
	}
}
