package retrieval

import (
	"github.com/KaramelBytes/docloom-cli/internal/utils"
	"strings"
)

// ChunkByTokens splits text into chunks of up to maxTokens, with overlap tokens between consecutive chunks.
// It uses a simple paragraph aggregator and token estimator for stability.
func ChunkByTokens(text string, maxTokens, overlap int) []string {
	if maxTokens <= 0 {
		maxTokens = 400
	}
	if overlap < 0 {
		overlap = 0
	}
	paras := splitParagraphs(text)
	var chunks []string
	var window []string
	curTokens := 0
	for _, p := range paras {
		t := utils.CountTokens(p)
		if curTokens+t > maxTokens && len(window) > 0 {
			chunks = append(chunks, strings.Join(window, "\n\n"))
			// prepare overlap
			if overlap > 0 {
				window, curTokens = backfillOverlap(window, overlap)
			} else {
				window = window[:0]
				curTokens = 0
			}
		}
		window = append(window, p)
		curTokens += t
	}
	if len(window) > 0 {
		chunks = append(chunks, strings.Join(window, "\n\n"))
	}
	return chunks
}

func splitParagraphs(s string) []string {
	raw := strings.Split(s, "\n\n")
	out := make([]string, 0, len(raw))
	for _, r := range raw {
		r = strings.TrimSpace(r)
		if r != "" {
			out = append(out, r)
		}
	}
	if len(out) == 0 && strings.TrimSpace(s) != "" {
		return []string{strings.TrimSpace(s)}
	}
	return out
}

func backfillOverlap(paras []string, overlap int) ([]string, int) {
	var out []string
	tokens := 0
	for i := len(paras) - 1; i >= 0; i-- {
		t := utils.CountTokens(paras[i])
		if tokens+t > overlap && len(out) > 0 {
			break
		}
		out = append([]string{paras[i]}, out...)
		tokens += t
	}
	return out, tokens
}
