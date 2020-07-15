# edgex-lf-edge-device-services

## Temperature Sensor 

Simulated MQTT device used in the Linux Foundation (LF Edge) blog. 

This process simulated a temperature sensor which publishes randomized temperature readings to `DataTopic` MQTT Topic. It also subscribes to `CommandTopic` for device commands from EdgeX Foundry publishing responses to `ResponseTopic`.

### Configuration

This process can be configured via CLI arguments:

- `--mqtt.host`: MQTT Broker host.
- `--mqtt.port`: MQTT Broker port.
- `--data.interval`: Interval in seconds at which the temperature values will publish to `DataTopic` topic.
- `--device`: Type of device.