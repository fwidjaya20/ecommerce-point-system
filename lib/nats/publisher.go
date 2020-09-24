package nats

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/nats-io/stan.go"
)

var globalPublisher *Publisher

type Publisher struct {
	publisher stan.Conn
	log       log.Logger
}

//StoreDetail store parameter
type StoreDetail struct {
	Channel     string
	Domain      string
	Subject     string
	EventSource string
	Data        interface{}
}

func NewPublisher(natsConn stan.Conn, logger *log.Logger) *Publisher {
	pub := &Publisher{
		publisher: natsConn,
	}

	if nil == logger {
		defaultLogger := log.NewLogfmtLogger(os.Stderr)
		defaultLogger = log.With(defaultLogger, "ts", log.DefaultTimestampUTC)
		defaultLogger = log.With(defaultLogger, "caller", log.DefaultCaller)
		defaultLogger = log.With(defaultLogger, "actor", "DefaultPublisher")

		pub.log = defaultLogger
	} else {
		pub.log = *logger
	}

	return pub
}

// SetGlobalPublisher: set the Publisher into Global Singleton
func SetGlobalPublisher(publisher *Publisher) {
	globalPublisher = publisher
}

// GetGlobalPublisher: get the Global Publisher
func GetGlobalPublisher() *Publisher {
	return globalPublisher
}

func (p *Publisher) Store(info StoreDetail) error {
	var t bytes.Buffer
	var result = make(map[string]interface{}, 0)

	result["domain"] = info.Domain
	result["subject"] = info.Subject
	result["event_source"] = info.EventSource
	result["data"] = info.Data

	resultByte, err := json.Marshal(result)

	if nil != err {
		_ = p.log.Log("__METHOD_", "Store", "error_publish_commit", err)
	}

	t.Write(resultByte)

	err = p.publisher.Publish(info.Subject, t.Bytes())

	if nil == err {
		_ = p.log.Log("__METHOD__", "Store", "transport", "NATS", "__PUBLISHED_MESSAGE_ON__", info.Channel, "__WITH_SUBJECT__", info.Subject)
	} else {
		_ = p.log.Log("__METHOD__", "Store", "transport", "NATS", "__ERROR_PUBLISHING_MESSAGE_ON__", info.Channel, "__WITH_SUBJECT__", info.Subject)
	}

	return err
}
