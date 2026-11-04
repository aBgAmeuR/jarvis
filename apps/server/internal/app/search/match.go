package search

import (
	"slices"
	"strings"

	"jarvis/server/internal/domain/search"
)

type scored struct {
	hit   search.Hit
	score int
}

func matchAll(hits []search.Hit, q string) []scored {
	var matches []scored
	for _, hit := range hits {
		if score, ok := match(hit, q); ok {
			matches = append(matches, scored{hit: hit, score: score})
		}
	}
	return matches
}

// match scores a hit, lower is better. The name ranks first, then the id and the detail:
// whole text, prefix, word prefix, anywhere.
func match(hit search.Hit, q string) (int, bool) {
	if score, ok := matchText(hit.Name, q); ok {
		return score, true
	}
	id, idOK := matchText(hit.ID, q)
	detail, detailOK := matchText(hit.Detail, q)
	switch {
	case idOK && detailOK:
		return 4 + min(id, detail), true
	case idOK:
		return 4 + id, true
	case detailOK:
		return 4 + detail, true
	}
	return 0, false
}

func matchText(text, q string) (int, bool) {
	text = strings.ToLower(text)
	switch {
	case text == q:
		return 0, true
	case strings.HasPrefix(text, q):
		return 1, true
	case slices.ContainsFunc(strings.FieldsFunc(text, isSeparator), func(word string) bool { return strings.HasPrefix(word, q) }):
		return 2, true
	case strings.Contains(text, q):
		return 3, true
	}
	return 0, false
}

func isSeparator(r rune) bool {
	return strings.ContainsRune(" -_./:", r)
}
