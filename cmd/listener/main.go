package main

import (
	"encoding/json"
	"game-node-sync-hltb/internal/queue"
	"game-node-sync-hltb/internal/util"
	"game-node-sync-hltb/internal/util/redis"
	"github.com/hibiken/asynq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {
	rabbitMqUrl := util.RMQUrl()

	redisAddr := redis.Url()
	asyncqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})

	defer asyncqClient.Close()

	conn, err := amqp.Dial(rabbitMqUrl)

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return
	}

	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	err = channel.ExchangeDeclare("sync-hltb", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare("sync.hltb.update.request", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind("sync.hltb.update.request", "sync.hltb.update.request", "sync-hltb", false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := channel.Consume("sync.hltb.update.request", "", true, false, false, false, nil)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			request := queue.UpdateRequest{}
			err := json.Unmarshal(d.Body, &request)
			if err != nil {
				log.Printf(" [!] Failed to parse message: %s - error: %s", d.Body, err)
				continue
			}
			log.Printf(" [x] Received request to update gameId: %d - with name: %s", request.GameId, request.Name)

			task, err := queue.CreateUpdateTask(&request)
			if err != nil {
				log.Printf(" [!] Failed to create task for update request: %s", err)
				continue
			}

			taskInfo, err := asyncqClient.Enqueue(task, asynq.MaxRetry(2))
			if err != nil {
				log.Printf(" [!] Failed to enqueue update request: %s", err)
				continue
			}

			log.Printf(" [x] Enqueued task with id: %s", taskInfo.ID)
		}
	}()

	log.Printf(" [*] Waiting for messages on %s -> %s. To exit press CTRL+C", "sync-hltb", "sync.hltb.update.request")
	<-forever

}
