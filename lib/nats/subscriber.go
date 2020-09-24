package nats

import (
	"fmt"
	"github.com/go-kit/kit/log"
	"github.com/nats-io/stan.go"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type Subscriber struct {
	subscriber stan.Conn
	data       SubscriberData
	handler    stan.MsgHandler
	log        log.Logger
}

type SubscriberData struct {
	Subject    string
	QueueGroup string
	Durable    string
	StartAt    string
}

func NewSubscriber(
	conn stan.Conn,
	logger log.Logger,
	data SubscriberData,
	handler stan.MsgHandler,
) *Subscriber {
	sub := &Subscriber{
		subscriber: conn,
		handler:    handler,
		log:        logger,
		data:       data,
	}

	if reflect.ValueOf(logger).IsNil() {
		defaultLogger := log.NewLogfmtLogger(os.Stderr)
		defaultLogger = log.With(logger, "ts", log.DefaultTimestampUTC)
		defaultLogger = log.With(logger, "caller", log.DefaultCaller)
		defaultLogger = log.With(logger, "actor", "DefaultSubscriber")

		sub.log = defaultLogger
	}

	return sub
}

func (s *Subscriber) Subscribe() stan.Subscription {
	var opts []stan.SubscriptionOption

	if "" != s.data.Durable {
		opts = append(opts, stan.DurableName(s.data.Durable))
	}

	startOpt := s.getOption()
	if nil != startOpt {
		opts = append(opts, *startOpt)
	}

	sub, err := s.subscriber.QueueSubscribe(s.data.Subject, s.data.QueueGroup, s.handler, opts...)

	if err != nil {
		_ = s.log.Log("NATS", fmt.Sprintf("Error when subscribing topic %s", s.data.Subject))
		return nil
	}

	_ = s.log.Log("NATS", fmt.Sprintf("Subscribed topic %s with durable %s and start option %s", s.data.Subject, s.data.Durable, s.data.StartAt))
	return sub
}


func (s *Subscriber) getOption() *stan.SubscriptionOption {
	var startOpt stan.SubscriptionOption
	if s.data.StartAt == "all" {
		startOpt = stan.DeliverAllAvailable()
	} else if strings.Index(s.data.StartAt, "since:") == 0 {
		var option = strings.Split(s.data.StartAt, ":")
		ago, err := time.ParseDuration(option[1])
		if err != nil {
			_ = s.log.Log("NATS", fmt.Sprintf("Error when subscribing topic %s", s.data.Subject))
			_ = s.log.Log("err", err)
			return nil
		}
		startOpt = stan.StartAtTimeDelta(ago)
	} else if strings.Index(s.data.StartAt, "time:") == 0 {
		var option = strings.Split(s.data.StartAt, ":")
		intTimestamp, err := strconv.ParseInt(option[1], 10, 64)
		if err != nil {
			_ = s.log.Log("NATS", fmt.Sprintf("Error when subscribing topic %s", s.data.Subject))
			_ = s.log.Log("err", err)
			return nil
		}
		startOpt = stan.StartAtTime(time.Unix(intTimestamp, 0))
	} else if strings.Index(s.data.StartAt, "seqno:") == 0 {
		var option = strings.Split(s.data.StartAt, ":")
		intSeq, err := strconv.ParseUint(option[1], 10, 64)
		if err != nil {
			_ = s.log.Log("NATS", fmt.Sprintf("Error when subscribing topic %s", s.data.Subject))
			_ = s.log.Log("err", err)
			return nil
		}
		startOpt = stan.StartAtSequence(intSeq)
	} else if s.data.StartAt == "none" {
		return nil
	} else {
		panic("invalid start option")
	}
	return &startOpt
}