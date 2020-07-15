package temperature_sensor

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"gitlab.com/bonafide-technology/edgex-lf-device-services/internal/mqtt/publisher"
	"math/rand"
	"strconv"
)

type device struct {
	name             string
	temperatureValue float32
}

func New() *device {
	return &device{
		name: "TemperatureSensor",
	}
}

func (d *device) PublishData(topic string, p publisher.Publisher) {
	d.temperatureValue = 50 + rand.Float32()*(120-50)
	type data struct {
		Name        string `json:"name"`
		Temperature string `json:"temperature"`
		Cmd         string `json:"cmd"`
		Method      string `json:"method"`
	}

	b, _ := json.Marshal(&data{
		Name:        d.name,
		Temperature: fmt.Sprintf("%f", d.temperatureValue),
		Cmd:         "temperature",
		Method:      "get",
	})

	if err := p.Publish(topic, 0, false, b); err != nil {
		fmt.Println("Publish ", topic, "error: ", err.Error())
	}
}

func (d *device) PublishResponse(uuid string, topic string, p publisher.Publisher) {
	type response struct {
		Uuid        string `json:"uuid"`
		Temperature string `json:"temperature"`
	}

	b, _ := json.Marshal(&response{
		Uuid:        uuid,
		Temperature: fmt.Sprintf("%f", d.temperatureValue),
	})

	if err := p.Publish(topic, 0, false, b); err != nil {
		fmt.Println("Publish ", topic, "error: ", err.Error())
	}
}

func (d *device) Update(m mqtt.Message) {
	type message struct {
		Cmd         string `json:"cmd"`
		Method      string `json:"method"`
		Uuid        string `json:"uuid"`
		Temperature string `json:"temperature"`
	}

	var inbound *message
	if err := json.Unmarshal(m.Payload(), &inbound); err != nil {
		fmt.Println(err.Error())
	}

	f, _ := strconv.ParseFloat(inbound.Temperature, 32)
	d.temperatureValue = float32(f)
}
