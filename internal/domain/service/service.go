package service

import (
	"context"
	"l0/internal/domain/entity"
	"l0/internal/infrastructure/repository"
)

type Orders interface {
	GetOrdersByID(ctx context.Context, id string) (entity.Order, error)
	InsertOrder(ctx context.Context, w entity.Order) error
}

type Service struct {
	Orders Orders
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Orders: NewOrdersService(repos.Order),
	}
}
