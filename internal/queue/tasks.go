package queue

import (
	"encoding/json"
	"fmt"
	"game-node-sync-hltb/internal/search"
	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeUpdateRequest = "update.request"
)

type UpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateResponse struct {
	Id    int                     `json:"id"`
	Match search.HLTBResponseItem `json:"match"`
}

func CreateUpdateTask(r *UpdateRequest) (*asynq.Task, error) {
	payload, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	task := asynq.NewTask(
		TypeUpdateRequest, payload,
		asynq.TaskID(fmt.Sprintf("update-request-%d", r.Id)),
		asynq.MaxRetry(3),
	)

	return task, nil
}
