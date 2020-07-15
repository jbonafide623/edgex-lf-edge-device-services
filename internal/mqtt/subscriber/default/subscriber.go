package _default

import mqtt "github.com/eclipse/paho.mqtt.golang"

type subscriber struct {
	client mqtt.Client
}

func New(client mqtt.Client) *subscriber {
	return &subscriber{
		client: client,
	}
}

func (s *subscriber) Subscribe(topic string, qos byte, handler func(client mqtt.Client, message mqtt.Message)) error {
	if token := s.client.Subscribe(topic, qos, handler); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
