package main

import (
	"game-node-sync-hltb/internal/queue"
	"game-node-sync-hltb/internal/util"
	"game-node-sync-hltb/internal/util/redis"
	"github.com/hibiken/asynq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func main() {

	rabbitMqUrl := util.RMQUrl()
	conn, err := amqp.Dial(rabbitMqUrl)

	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
		return
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	err = channel.ExchangeDeclare("sync-hltb", "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	_, err = channel.QueueDeclare("sync.hltb.update.response", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.QueueBind("sync.hltb.update.response", "sync.hltb.update.response", "sync-hltb", false, nil)
	if err != nil {
		log.Fatal(err)
	}

	redisAddr := redis.Url()
	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: redisAddr},
		asynq.Config{
			// Specify how many concurrent workers to use
			Concurrency: 1,
			// See the godoc for other configuration options
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(queue.TypeUpdateRequest, queue.HandleUpdateRequest)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
