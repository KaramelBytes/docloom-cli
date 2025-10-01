package utils

// Simple token estimation utilities.
// These are placeholders and can be refined later to match specific model tokenization.

// CountTokens estimates the number of tokens in the given text.
// For MVP, we approximate 1 token ~= 4 characters (rough heuristic).
func CountTokens(text string) int {
	if len(text) == 0 {
		return 0
	}
	// Ensure at least 1 token for any non-empty text
	tokens := len([]rune(text)) / 4
	if tokens == 0 {
		return 1
	}
	return tokens
}

// TruncateToTokenLimit naively truncates text to roughly fit within a token limit.
func TruncateToTokenLimit(text string, limit int) string {
	if limit <= 0 {
		return ""
	}
	runes := []rune(text)
	// Expand limit to character count using the same 4 chars per token heuristic
	charLimit := limit * 4
	if charLimit >= len(runes) {
		return text
	}
	return string(runes[:charLimit])
}

// TokenBreakdown returns a simple breakdown map of labeled sections to token counts.
func TokenBreakdown(sections map[string]string) map[string]int {
	out := make(map[string]int, len(sections))
	for k, v := range sections {
		out[k] = CountTokens(v)
	}
	return out
}
