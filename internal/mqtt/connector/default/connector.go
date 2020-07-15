package _default

import mqtt "github.com/eclipse/paho.mqtt.golang"

type connector struct{}

func New() *connector {
	return &connector{}
}

func (c *connector) Connect(client mqtt.Client) error {
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	return nil
}
