package container

import (
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/repositories"
	"github.com/go-kit/kit/log"
)

type Container struct {
	UserPointService user_point.UseCase
}

func New(
	logger log.Logger,
) Container {
	return Container{
		UserPointService: user_point.NewUserPointService(logger, repositories.NewUserPointRepository()),
	}
}