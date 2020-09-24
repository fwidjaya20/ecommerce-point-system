package nats

import (
	"encoding/json"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/transports/nats/handlers/add"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/transports/nats/handlers/deduct"
	"github.com/fwidjaya20/ecommerce-point-system/lib/nats"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/nats-io/stan.go"
)

func NewSubscriber(conn stan.Conn, logger log.Logger, service user_point.UseCase) ([]stan.Subscription, error) {
	var subs []stan.Subscription
	var err error

	addUserPoint, err := addUserPointSubscriber(conn, logger, service)
	subs = append(subs, addUserPoint)

	deductUserPoint, err := deductUserPointSubscriber(conn, logger, service)
	subs = append(subs, deductUserPoint)

	return subs, err
}

func addUserPointSubscriber(conn stan.Conn, logger log.Logger, service user_point.UseCase) (stan.Subscription, error) {
	var err error

	addUserPoint := nats.NewSubscriber(conn, logger, nats.SubscriberData{
		Subject:    "user_point.add",
		QueueGroup: "user_point.queue",
		Durable:    "user_point_sub",
		StartAt:    "all",
	}, func(msg *stan.Msg) {
		var stream map[string]interface{}
		var err error

		if err := json.Unmarshal(msg.Data, &stream); nil != err {
			_ = logger.Log("err", err)
			panic("error")
		}

		stream = stream["data"].(map[string]interface{})

		versionHandler := add.MessageVersion{
			UserPointService: service,
		}

		if stream["data"] != nil {
			versionHandler.Data = stream["data"].(interface{})
		}

		if stream["metadata"] != nil {
			versionHandler.Metadata = stream["metadata"].(map[string]interface{})
		}

		if err = versionHandler.Handle(); nil != err {
			_ = level.Error(logger).Log("error", err)
			panic(err)
		}
	})

	return addUserPoint.Subscribe(), err
}

func deductUserPointSubscriber(conn stan.Conn, logger log.Logger, service user_point.UseCase) (stan.Subscription, error) {
	var err error

	addUserPoint := nats.NewSubscriber(conn, logger, nats.SubscriberData{
		Subject:    "user_point.deduct",
		QueueGroup: "user_point.queue",
		Durable:    "user_point_sub",
		StartAt:    "all",
	}, func(msg *stan.Msg) {
		var stream map[string]interface{}
		var err error

		if err := json.Unmarshal(msg.Data, &stream); nil != err {
			_ = logger.Log("err", err)
			panic("error")
		}

		stream = stream["data"].(map[string]interface{})

		versionHandler := deduct.MessageVersion{
			UserPointService: service,
		}

		if stream["data"] != nil {
			versionHandler.Data = stream["data"].(interface{})
		}

		if stream["metadata"] != nil {
			versionHandler.Metadata = stream["metadata"].(map[string]interface{})
		}

		if err = versionHandler.Handle(); nil != err {
			_ = level.Error(logger).Log("error", err)
			panic(err)
		}
	})

	return addUserPoint.Subscribe(), err
}
