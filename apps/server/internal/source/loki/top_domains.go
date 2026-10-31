package loki

import (
	"cmp"
	"context"
	"slices"
	"strconv"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/network"
	"jarvis/server/internal/source/failure"
)

func (c *Client) topDomains(ctx context.Context, selector string, status network.DomainStatus, w domain.Window, limit int) ([]network.DomainCount, error) {
	response := `| response_type!="BLOCKED"`
	if status == network.DomainBlocked {
		response = `| response_type="BLOCKED"`
	}
	span := strconv.FormatInt(int64(w.End.Sub(w.Start).Seconds()), 10)
	expr := "topk(" + strconv.Itoa(limit) + ", sum by (question_name) (count_over_time(" + selector + " " + response + " [" + span + "s])))"
	entries, err := c.queryVector(ctx, expr, w.End)
	if err != nil {
		return nil, failure.Wrap(name, err)
	}
	domains := make([]network.DomainCount, 0, len(entries))
	for _, e := range entries {
		domains = append(domains, network.DomainCount{Domain: e.labels["question_name"], Count: int64(e.value)})
	}
	slices.SortFunc(domains, func(a, b network.DomainCount) int { return cmp.Compare(b.Count, a.Count) })
	return domains, nil
}
