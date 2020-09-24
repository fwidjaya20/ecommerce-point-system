package repositories

import (
	"context"
	"github.com/fwidjaya20/ecommerce-point-system/internal/databases/models"
)

type Interface interface {
	StoreEvent(ctx context.Context, model *models.UserPointEvent) error
	GetPoint(ctx context.Context, userId string) (*models.UserPointSnapshot, error)
}