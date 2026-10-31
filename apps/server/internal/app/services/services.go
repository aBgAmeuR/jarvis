package services

import "time"

type Services struct {
	catalog CatalogSource
	signals SignalSource
	now     func() time.Time
}

func NewServices(catalog CatalogSource, signals SignalSource) *Services {
	return &Services{catalog: catalog, signals: signals, now: time.Now}
}
