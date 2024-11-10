package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// Try to connect to rabbitmq
	rabbitConn, err := connet()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ!")
	// Start listening for messages from rabbitmq

	// create a consumer

	// watch the queue and consume events

}

func connet() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not ready yet ready...")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println("Unable to connect to RabbitMQ, giving up.")
			return nil, err
		}

		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off...")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}