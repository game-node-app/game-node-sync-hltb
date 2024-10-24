package scraper

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func GetApiKey() (string, error) {
	hltbBaseUrl := "https://howlongtobeat.com"

	// Colly collector
	c := colly.NewCollector(
		colly.Async(true),
	)

	var result string
	var err error = nil

	err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Delay:       1 * time.Second,
		RandomDelay: 1 * time.Second, // Add some randomness to the delay
	})

	// Set Fake User Agent
	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36")
	})

	c.OnHTML("script", func(e *colly.HTMLElement) {
		var src = e.Attr("src")
		var targetScript string
		if !strings.HasPrefix(src, "/_next/static/chunks/pages/_app") {
			return
		}

		targetScript = src

		fmt.Println(fmt.Sprintf("https://howlongtobeat.com%s", targetScript))
		request, err := http.NewRequest("GET", fmt.Sprintf("https://howlongtobeat.com%s", targetScript), nil)
		request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Mobile Safari/537.36")
		if err != nil {
			log.Fatal(err)
			return
		}
		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Fatal(err)
			return
		}

		bodyBytes, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		bodyStr := string(bodyBytes)
		//fmt.Println(bodyStr)
		r := regexp.MustCompile("fetch\\(\"/api/search/\"\\.concat\\(\"([^\"]+)\"\\)")
		submatch := r.FindStringSubmatch(bodyStr)
		if len(submatch) > 0 {
			result = submatch[1]
		} else {
			err = errors.New("could not find API key submatch")
		}
	})

	err = c.Visit(hltbBaseUrl)

	c.Wait()

	return result, err

}
