package search

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	originalName := "The.Witcher.3 GOTY Edition     "
	parsedName := parseGameName(originalName)
	fmt.Println(parsedName)
	assert.Equal(t, "the witcher 3", parsedName)
}
