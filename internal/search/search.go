package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"game-node-sync-hltb/internal/scraper"
	"log"
	"net/http"
	"strings"
)

func searchBody(searchCriteria []string) *HLTBSearchRequest {
	return &HLTBSearchRequest{
		SearchType:  "games",
		SearchTerms: searchCriteria,
		SearchPage:  1,
		Size:        20,
	}
}

func searchEndpoint(apiKey string) string {
	return fmt.Sprintf("https://howlongtobeat.com/api/search/%s", apiKey)
}

func Games(q string) (*HLTBResponse, error) {
	apiKey, err := scraper.GetApiKey()
	if err != nil {
		return nil, err
	}

	parsedGameName := parseGameName(q)
	log.Printf(" [X] Parsed game game - from: %s to: %s", q, parsedGameName)
	searchCriteria := strings.Split(parsedGameName, " ")
	targetUrl := searchEndpoint(apiKey)
	body := searchBody(searchCriteria)

	bodyJson, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", targetUrl, bytes.NewReader(bodyJson))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	request.Header.Set("Origin", "https://howlongtobeat.com")
	request.Header.Set("Referer", "https://howlongtobeat.com")

	if err != nil {
		return nil, err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	hltbResponse := HLTBResponse{}
	err = json.NewDecoder(response.Body).Decode(&hltbResponse)
	if err != nil {
		return nil, err
	}

	return &hltbResponse, nil

}
