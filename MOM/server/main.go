package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var config = GetConfig()

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@" + config.IP + ":" + config.PORT + "/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}

	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		for d := range msgs {

			info := string(d.Body)
			var response string

			if info == "list" {

				response = listFiles()
			} else {
				response = strconv.FormatBool(searchFile(info))
			}

			err = ch.PublishWithContext(ctx,
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(response),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func searchFile(name string) bool {

	if _, err := os.Stat(config.DIR + name); err == nil {
		return true
	} else {
		return false
	}
}

func listFiles() string {

	names := make([]string, 0)

	files, err := ioutil.ReadDir(config.DIR)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		names = append(names, file.Name())
	}

	return strings.Join(names[:], " , ")
}
