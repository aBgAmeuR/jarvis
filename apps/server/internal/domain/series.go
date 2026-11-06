package domain

import "time"

type Unit string

const (
	UnitRatio          Unit = "ratio"
	UnitBytes          Unit = "bytes"
	UnitBytesPerSecond Unit = "bytesPerSecond"
	UnitLoad           Unit = "load"
	UnitCelsius        Unit = "celsius"
	UnitSeconds        Unit = "seconds"
	UnitPerSecond      Unit = "perSecond"
)

// A gap is a missing timestamp.
type Series struct {
	Name string
	Unit Unit
	// Query is the PromQL that produced the series, for Grafana Explore.
	Query      string
	Timestamps []time.Time
	Values     []float64
	Summary    *Summary
}

type Summary struct {
	Avg  float64
	Peak float64
}

type TimeSeries struct {
	Window Window
	Series []Series
}
