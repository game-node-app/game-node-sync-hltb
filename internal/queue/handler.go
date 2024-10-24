package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"game-node-sync-hltb/internal/search"
	"github.com/hibiken/asynq"
	"log"
)

func HandleUpdateRequest(ctx context.Context, t *asynq.Task) error {
	var r UpdateRequest

	err := json.Unmarshal(t.Payload(), &r)
	if err != nil {
		return err
	}

	//id := r.Id
	name := r.Name

	hltbResp, err := search.Games(name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hltbResp)

	return nil

}
