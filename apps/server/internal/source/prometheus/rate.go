package prometheus

import (
	"strconv"

	"jarvis/server/internal/domain"
)

func rateRange(w domain.Window) string {
	return "[" + strconv.FormatInt(int64(w.RateInterval().Seconds()), 10) + "s]"
}
