package search

import (
	"regexp"
	"strings"
)

func parseGameName(name string) string {

	releaseYearRegex := regexp.MustCompile(`\(\d{4}\)`)
	specialEditionRegex := regexp.MustCompile(`(?i)(The |Digital )?(GOTY|Deluxe|Standard|Ultimate|Definitive|Enhanced|Collector's|Premium|Digital|Limited|Game of the Year|Special|Reloaded|DIRECTOR'S CUT|\d{4}) Edition`)
	underscoreRegex := regexp.MustCompile(`_`)
	dotsRegex := regexp.MustCompile(`\.`)
	specialCharsRegex := regexp.MustCompile(`[^A-Za-z0-9 ]`)
	duplicateSpacesRegex := regexp.MustCompile(`\s{2,}`)

	result := name
	result = strings.ToLower(result)
	result = releaseYearRegex.ReplaceAllString(result, "")
	result = specialEditionRegex.ReplaceAllString(result, "")
	result = underscoreRegex.ReplaceAllString(result, " ")
	result = dotsRegex.ReplaceAllString(result, " ")
	result = specialCharsRegex.ReplaceAllString(result, "")
	result = duplicateSpacesRegex.ReplaceAllString(result, " ")
	result = strings.TrimSpace(result)

	return result
}
