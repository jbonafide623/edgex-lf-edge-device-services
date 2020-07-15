package _default

import mqtt "github.com/eclipse/paho.mqtt.golang"

type publisher struct {
	client mqtt.Client
}

func New(client mqtt.Client) *publisher {
	return &publisher{
		client: client,
	}
}

func (p *publisher) Publish(topic string, qos byte, retain bool, payload interface{}) error {
	if token := p.client.Publish(topic, qos, retain, payload); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
