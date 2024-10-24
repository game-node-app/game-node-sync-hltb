package main

import (
	"encoding/json"
	"fmt"
	"game-node-sync-hltb/internal/search"
	"log"
)

func main() {

	//rabbitMqUrl := util.GetEnv("RABBITMQ_URL", "amqp://gamenode:gamenode@localhost:5672")
	//
	//connection, err := amqp.Dial(rabbitMqUrl)
	//if err != nil {
	//	log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	//	return
	//}
	//
	//err = queue.StartListening(connection)
	//if err != nil {
	//	log.Fatalf("Failed to listen in RabbitMQ: %s", err)
	//	return
	//}
	//
	//defer connection.Close()

	response, err := search.Games("Witcher 3")
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := json.MarshalIndent(response, "", "\t")
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
