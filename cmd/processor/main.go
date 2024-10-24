package main

import (
	"game-node-sync-hltb/internal/queue"
	"game-node-sync-hltb/internal/util"
	"github.com/hibiken/asynq"
	"log"
)

func main() {
	redisAddr := util.RedisURL()
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

	log.Printf(" [X] Starting worker processor server...")
	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}

}
