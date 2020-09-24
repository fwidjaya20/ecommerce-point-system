package user_point

import (
	"context"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/models"
)

type UseCase interface {
	PublishAddPoint(ctx context.Context, payload models.PublishAddOrDeductPointRequest) (*models.UserPointInquiryResponse, error)
	PublishDeductPoint(ctx context.Context, payload models.PublishAddOrDeductPointRequest) (*models.UserPointInquiryResponse, error)
	StoreEvent(ctx context.Context, payload models.AddOrDeductPointRequest) error
	GetPoint(ctx context.Context, payload models.GetPointRequest) (*models.UserPointResponse, error)
}
