package services

import (
	"context"

	"jarvis/server/internal/domain"
	"jarvis/server/internal/domain/services"
)

func (s *Services) Get(ctx context.Context, slug string) (services.Service, error) {
	list, err := s.List(ctx)
	if err != nil {
		return services.Service{}, err
	}
	for _, service := range list {
		if service.Slug == slug {
			return service, nil
		}
	}
	return services.Service{}, domain.NewNotFound("service", slug)
}
