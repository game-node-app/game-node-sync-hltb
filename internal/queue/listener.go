package queue

import (
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func StartListening(conn *amqp.Connection) error {
	channel, err := conn.Channel()
	if err != nil {
		return err
	}

	_, err = channel.QueueDeclare("sync-hltb", true, false, false, false, nil)
	if err != nil {
		return err
	}

	err = channel.QueueBind("sync-hltb", "update.request", "sync", false, nil)
	if err != nil {
		return err
	}

	err = channel.QueueBind("sync-hltb", "update.response", "sync", false, nil)

	msgs, err := channel.Consume("sync-hltb", "", true, false, false, false, nil)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			request := UpdateRequest{}
			err := json.Unmarshal(d.Body, &request)
			if err != nil {
				log.Fatalf("Failed to parse message: %s - error: %s", d.Body, err)
				return
			}

			log.Printf(" [x] Received request to update gameId: %d - with name: %s", request.Id, request.Name)
		}
	}()

	log.Printf(" [*] Waiting for messages on %s -> %s. To exit press CTRL+C", "sync-hltb", "update.request")
	<-forever

	return nil
}
