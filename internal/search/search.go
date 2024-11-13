package search

import (
	"bytes"
	"encoding/json"
	"fmt"
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
		SearchOptions: HLTBSearchRequestOptions{
			Users: HLTBSearchRequestOptionsUsers{
				Id:           "90f8120e015db09f",
				SortCategory: "postcount",
			},
		},
	}
}

func searchEndpoint() string {
	return fmt.Sprintf("https://howlongtobeat.com/api/search")
}

func Games(q string) (*HLTBResponse, error) {
	// API Key temporarily not necessary...
	//apiKey, err := scraper.GetApiKey()
	//if err != nil {
	//	log.Printf(" [!] Failed to retrieve API Key: %v", err)
	//	return nil, err
	//}

	parsedGameName := parseGameName(q)
	log.Printf(" [x] Parsed name - from: %s to: %s", q, parsedGameName)
	searchCriteria := strings.Split(parsedGameName, " ")
	targetUrl := searchEndpoint()
	body := searchBody(searchCriteria)

	bodyJson, err := json.Marshal(body)
	if err != nil {
		log.Printf(" [!] Failed to marshal request body (send pending): %v", err)
		return nil, err
	}

	request, err := http.NewRequest("POST", targetUrl, bytes.NewReader(bodyJson))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	request.Header.Set("Origin", "https://howlongtobeat.com")
	request.Header.Set("Referer", "https://howlongtobeat.com")

	if err != nil {
		log.Printf(" [!] Building Request to HLTB API failed: %v", err)
		return nil, err
	}

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		log.Printf(" [!] Sending request to HLTB API failed: %v", err)
		return nil, err
	}

	hltbResponse := HLTBResponse{}
	err = json.NewDecoder(response.Body).Decode(&hltbResponse)
	if err != nil {
		log.Printf(" [!] Failed to Marshal HLTB Response: %v", err)
		return nil, err
	}

	return &hltbResponse, nil

}
