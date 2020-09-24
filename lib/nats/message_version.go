package nats

type MessageVersion interface {
	Handle(msg interface{}, metadata interface{}) error
}
