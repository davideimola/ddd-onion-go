package order

import (
	"context"
	"github.com/google/uuid"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *Order) error
	UpdateOrder(ctx context.Context, orderID uuid.UUID, updateFn func(o *Order) error) (*Order, error)
}
