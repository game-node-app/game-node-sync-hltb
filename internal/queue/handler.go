package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"game-node-sync-hltb/internal/search"
	"game-node-sync-hltb/internal/util"
	"game-node-sync-hltb/internal/util/redis"
	"github.com/hibiken/asynq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

const (
	FailedAttemptStoreKey = "hltb-failed-attempt-%d"
	SuccessStoreKey       = "hltb-success-%d"
)

// Should only be used for matchless tries.
func storeFailedAttempt(gameId int) {
	var duration = 3 * 24 * time.Hour
	var key = fmt.Sprintf(FailedAttemptStoreKey, gameId)
	err := redis.Set(key, "true", &duration)
	if err != nil {
		log.Printf("failed to store failed attempt %d: %v", gameId, err)
	}
}

func storeSuccess(gameId int) {
	var duration = 7 * 24 * time.Hour
	var key = fmt.Sprintf(SuccessStoreKey, gameId)
	err := redis.Set(key, "true", &duration)
	if err != nil {
		log.Printf("failed to store success attempt %d: %v", gameId, err)
	}
}

func hasFailedAttempt(gameId int) bool {
	var key = fmt.Sprintf(FailedAttemptStoreKey, gameId)
	result, err := redis.Get(key)
	if err != nil {
		return false
	}

	return result != ""
}

func hasSuccess(gameId int) bool {
	var key = fmt.Sprintf(SuccessStoreKey, gameId)
	result, err := redis.Get(key)
	if err != nil {
		return false
	}

	return result != ""
}

func publishMatch(res *UpdateResponse) error {
	rabbitMqUrl := util.RMQUrl()
	conn, err := amqp.Dial(rabbitMqUrl)

	if err != nil {
		log.Printf("Failed to connect to RabbitMQ: %s", err)
		return err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return err
	}

	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("Failed to marshal response error: %s", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = channel.PublishWithContext(ctx, "sync-hltb", "sync.hltb.update.response", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        resBytes,
	})
	if err != nil {
		log.Printf(" [!] Failed to publish message! %s", err)
	}

	log.Printf(" [x] Sucessfully published response to queue: %s", "sync.hltb.update.response")

	return nil
}

func HandleUpdateRequest(ctx context.Context, t *asynq.Task) error {
	var r UpdateRequest

	err := json.Unmarshal(t.Payload(), &r)
	if err != nil {
		return err
	}

	id := r.GameId
	name := r.Name

	if recentFail := hasFailedAttempt(id); recentFail {
		log.Printf(" [x] Skipping %d since it is recently failed...", id)
		return nil
	}

	if recentSuccess := hasSuccess(id); recentSuccess {
		log.Printf(" [x] Skipping %d since it has been processed recently...", id)
		return nil
	}

	hltbResp, err := search.Games(name)
	if err != nil {
		log.Printf(" [!] Failed to find matches for gameId: %d - error: %s", id, err)
		return fmt.Errorf("failed to find matches for gameId: %d - error: %s", id, err)
	}

	if hltbResp != nil && hltbResp.Data != nil && len(hltbResp.Data) > 0 {
		log.Printf(" [x] Found at least one match for gameId: %d", id)
		res := UpdateResponse{
			GameId: id,
			Match:  hltbResp.Data[0],
		}
		err = publishMatch(&res)
		if err != nil {
			return err
		}

		storeSuccess(id)
	} else {
		log.Printf(" [x] No match found for gameId: %d", id)
		storeFailedAttempt(id)
	}

	// Request was successful, even if no results were found
	time.Sleep(4 * time.Second)
	log.Printf(" [x] Finished processing request for gameId: %d", id)

	return nil
}
