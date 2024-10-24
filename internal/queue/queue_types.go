package queue

import "game-node-sync-hltb/internal/search"

type UpdateRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UpdateResponse struct {
	Id    int                     `json:"id"`
	Match search.HLTBResponseItem `json:"match"`
}
