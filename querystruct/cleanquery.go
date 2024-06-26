package querystruct

import (
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	cleanNewlineRegex = regexp.MustCompile(`\n+|\t+`)
	cleanSpaceRegex   = regexp.MustCompile(` +`)
)

// CleanQuery is helper to clean query generated using text/template. The helper clean extra spaces and new line.
func CleanQuery(query string) string {
	newlineCleared := cleanNewlineRegex.ReplaceAllString(query, " ")
	return strings.TrimSpace(cleanSpaceRegex.ReplaceAllString(newlineCleared, " "))
}

// AssertCleanQueryEqual is test helper to compare query generated using text/template. The helper clean extra spaces and new line then compare the query string with the expected query string.
func AssertCleanQueryEqual(t *testing.T, expected, actual string) {
	assert.Equal(t, CleanQuery(expected), CleanQuery(actual))
}
