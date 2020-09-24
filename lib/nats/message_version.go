package nats

type MessageVersionHandler interface {
	Handle(msg interface{}, metadata interface{}) error
}
