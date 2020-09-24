package nats

import (
	"github.com/fwidjaya20/ecommerce-point-system/cmd/container"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/transports/nats"
	"github.com/go-kit/kit/log"
	"github.com/nats-io/stan.go"
)

func NATSSubscribers(conn stan.Conn, containers container.Container, logger log.Logger) ([]stan.Subscription, error) {
	var subs []stan.Subscription
	var err error

	subUserPoint, err := nats.NewSubscriber(conn, logger, containers.UserPointService)
	if nil != err {
		return nil, err
	}
	subs = append(subs, subUserPoint...)

	return subs, err
}