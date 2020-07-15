package main

import (
	"encoding/json"
	"flag"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	deviceFactory "gitlab.com/bonafide-technology/edgex-lf-device-services/internal/device/factory/default"
	_default "gitlab.com/bonafide-technology/edgex-lf-device-services/internal/mqtt/connector/default"
	publisher "gitlab.com/bonafide-technology/edgex-lf-device-services/internal/mqtt/publisher/default"
	subscriber "gitlab.com/bonafide-technology/edgex-lf-device-services/internal/mqtt/subscriber/default"
	"os"
	"sync"
	"time"
)

const (
	commandTopic  = "CommandTopic"
	dataTopic     = "DataTopic"
	responseTopic = "ResponseTopic"
)

func main() {
	var brokerHost string
	var brokerPort string
	var dataInterval time.Duration
	var dev string
	flag.StringVar(&brokerHost, "mqtt.host", "localhost", "MQTT Broker Host")
	flag.StringVar(&brokerPort, "mqtt.port", "1883", "MQTT Broker Port")
	flag.DurationVar(&dataInterval, "data.interval", 5, "Publish data interval")
	flag.StringVar(&dev, "device", "temperature", "Device")
	flag.Parse()

	wg := &sync.WaitGroup{}
	options := mqtt.NewClientOptions()
	options.AddBroker(fmt.Sprintf("%s:%s", brokerHost, brokerPort))
	client := mqtt.NewClient(options)

	connector := _default.New()
	if err := connector.Connect(client); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	p := publisher.New(client)
	s := subscriber.New(client)
	d := deviceFactory.New().Get(dev)

	ticker := time.NewTicker(dataInterval * time.Second)
	quit := make(chan struct{})
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("DATA...")
				d.PublishData(dataTopic, p)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	if err := s.Subscribe(commandTopic, 0, func(client mqtt.Client, message mqtt.Message) {
		type Message struct {
			Cmd    string `json:"cmd"`
			Method string `json:"method"`
			Uuid   string `json:"uuid"`
		}

		var inbound *Message
		if unmarshalErr := json.Unmarshal(message.Payload(), &inbound); unmarshalErr != nil {
			fmt.Println("Subscribe command message error: ", unmarshalErr.Error())
		}

		if inbound.Method == "put" {
			d.Update(message)
		}

		// Response data
		d.PublishResponse(inbound.Uuid, responseTopic, p)

	}); err != nil {
		fmt.Println("Subscribe to ", commandTopic, " error: ", err.Error())
	}
	wg.Wait()
}
