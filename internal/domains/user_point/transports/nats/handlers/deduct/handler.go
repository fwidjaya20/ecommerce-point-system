package deduct

import (
	"errors"
	"fmt"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point"
	"github.com/fwidjaya20/ecommerce-point-system/lib/nats"
)

type MessageVersion struct {
	UserPointService user_point.UseCase
	Data             interface{}
	Metadata         map[string]interface{}
	handler          nats.MessageVersionHandler
}

func (mv *MessageVersion) Handle() error {
	if nil == mv.Metadata {
		return nil
	}

	switch mv.Metadata["version"].(float64) {
	case 1:
		mv.handler = UserPointHandler1{
			UserPointService: mv.UserPointService,
		}
	}

	if mv.handler == nil {
		return errors.New(fmt.Sprintf("[NATS-UserPoint-Deduct] there is no handler for version %s", mv.Metadata["version"]))
	}

	return mv.handler.Handle(mv.Data, mv.Metadata)
}
