package availability

type Silences struct {
	source SilenceSource
}

func NewSilences(source SilenceSource) *Silences {
	return &Silences{source: source}
}
