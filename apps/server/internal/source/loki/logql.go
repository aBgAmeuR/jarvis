package loki

import (
	"strconv"
	"strings"

	"jarvis/server/internal/domain/logs"
)

func (c *Client) LogQuery(q logs.Query) string { return logQL(q) }

// Loki wants at least one matcher, so a query with none selects every stream with a service name.
func logQL(q logs.Query) string {
	var matchers []string
	if q.VMID != 0 {
		matchers = append(matchers, "vmid="+strconv.Quote(strconv.Itoa(q.VMID)))
	}
	if q.Service != "" {
		matchers = append(matchers, "service_name="+strconv.Quote(q.Service))
	}
	if q.Level != "" {
		matchers = append(matchers, "level="+strconv.Quote(q.Level))
	}
	if len(matchers) == 0 {
		matchers = append(matchers, `service_name=~".+"`)
	}
	expr := "{" + strings.Join(matchers, ",") + "}"
	if q.Contains != "" {
		expr += " |= " + strconv.Quote(q.Contains)
	}
	return expr
}
