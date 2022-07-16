package repository

import (
	"context"
	"l0/internal/domain/entity"

	"github.com/gocraft/dbr"
)

type DBConn interface {
	BeginTx(ctx context.Context, f func(tx *dbr.Tx) error) error
}
type Order interface {
	GetOrdersByID(ctx context.Context, id string) (entity.Order, error)
	InsertOrder(ctx context.Context, w entity.Order) error
}

type Repository struct {
	Order
}

type baseConn struct {
	*dbr.Connection
}

func (r *baseConn) BeginTx(ctx context.Context, f func(tx *dbr.Tx) error) error {
	tx, err := r.NewSession(nil).BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer tx.RollbackUnlessCommitted()

	err = f(tx)
	if err != nil {
		return err
	}

	if tx.Commit() != nil {
		return err
	}

	return nil
}

func NewRepository(db *dbr.Connection) *Repository {
	base := &baseConn{db}

	return &Repository{
		Order: NewOrdersRepository(base),
	}
}
