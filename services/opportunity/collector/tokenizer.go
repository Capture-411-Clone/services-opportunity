package collector

import (
	"strings"
	"unicode/utf8"

	"github.com/microcosm-cc/bluemonday"
)

// SimpleTokenize approximates tokenization by splitting text into words and trims spaces and newlines.
func SimpleTokenize(text string) []string {
	// Sanitize to remove HTML tags and trim spaces and newlines.
	p := bluemonday.StrictPolicy()
	sanitizedText := strings.TrimSpace(p.Sanitize(text))

	// Split text into words based on spaces after trimming.
	words := strings.Fields(sanitizedText)
	return words
}

// ShrinkTextToFitLimit attempts to reduce the text to fit within the token limit.
// It returns the adjusted text, a boolean indicating whether the original text was within the limit, and the token count.
func ShrinkTextToFitLimit(text string, limit int) (string, bool, int) {
	tokens := SimpleTokenize(text)
	tokenCount := 0
	var resultTokens []string

	for _, token := range tokens {
		newTokenCount := tokenCount + utf8.RuneCountInString(token)
		if newTokenCount > limit {
			break // Stop adding tokens if the next token would exceed the limit.
		}
		resultTokens = append(resultTokens, token)
		tokenCount = newTokenCount
	}

	// Reconstruct the text from the tokens that fit within the limit.
	shrunkText := strings.Join(resultTokens, " ")
	withinLimit := len(tokens) == len(resultTokens)

	return shrunkText, withinLimit, tokenCount
}

// TokenLimitCheck checks if the tokenized input exceeds the token limit.
func TokenLimitCheck(text string, limit int) (bool, int) {
	tokens := SimpleTokenize(text)
	tokenCount := 0

	for _, token := range tokens {
		// Count runes in each token.
		tokenCount += utf8.RuneCountInString(token)
	}

	// Check if token count exceeds the limit.
	return tokenCount > limit, tokenCount
}
