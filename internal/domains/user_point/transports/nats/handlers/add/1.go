package add

import (
	"context"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/endpoints"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/models"
)

type UserPointHandler1 struct {
	UserPointService user_point.UseCase
}

func (h UserPointHandler1) Handle(
	data interface{},
	metadata interface{},
) error {
	endpoint := endpoints.StoreEvent(h.UserPointService)

	var d = data.(map[string]interface{})

	var payload = models.AddOrDeductPointRequest{
		Id:        d["Id"].(string),
		UserId:    d["UserId"].(string),
		Point:     d["Point"].(float64),
		PointType: d["PointType"].(string),
		Notes:     d["Notes"].(string),
	}

	_, err := endpoint(context.Background(), &payload)

	if nil != err {
		return err
	}

	return nil
}
