package service

import (
	"context"
	"l0/internal/domain/entity"
	"l0/internal/infrastructure/repository"
)

type OrdersService struct {
	repo repository.Order
}

func NewOrdersService(repo repository.Order) *OrdersService {
	return &OrdersService{repo: repo}
}

func (s *OrdersService) GetOrdersByID(ctx context.Context, id string) (entity.Order, error) {
	return s.repo.GetOrdersByID(ctx, id)
}
func (s *OrdersService) InsertOrder(ctx context.Context, w entity.Order) error {
	return s.repo.InsertOrder(ctx, w)
}
