package main

import mqttClient "airport-weather/internal/mqtt-client"

func main() {
	client := mqttClient.GetMqttClient("temperature-sensor")
	//TODO: We want to be able to send those data continuously, like every one second with different values
	mqttClient.Publish(client, "{\"sensorId\":\"2\",\"airportId\":\"CDG\",\"sensorType\":\"temperature\",\"value\":\"20\",\"timestamp\":\"20060102150405\"}}")
}
