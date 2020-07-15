package publisher

type Publisher interface {
	Publish(topic string, qos byte, retain bool, payload interface{}) error
}
