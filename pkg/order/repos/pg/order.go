package pgOrderRepository

import (
	"context"
	apperrors "davideimola.dev/ddd-onion/internal/errors/app"
	"davideimola.dev/ddd-onion/pkg/order"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type OrderRepository struct {
	db *pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) order.OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order *order.Order) error {
	q := New(r.db)

	return q.CreateOrder(ctx, CreateOrderParams{
		ID:         order.ID,
		Status:     order.Status(),
		CreatedAt:  pgtype.Timestamp{Time: order.CreatedAt, Valid: true},
		UpdatedAt:  pgtype.Timestamp{Time: order.UpdatedAt, Valid: true},
		CustomerID: order.CustomerID,
	})
}

func (r *OrderRepository) UpdateOrder(ctx context.Context, orderID uuid.UUID, updateFn func(*order.Order) error) (*order.Order, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	q := New(tx)

	o, err := r.getOrderForUpdate(ctx, q, orderID)

	err = updateFn(o)
	if err != nil {
		return nil, err
	}

	err = q.UpdateOrder(ctx, UpdateOrderParams{
		OrderID:   o.ID,
		Status:    o.Status(),
		UpdatedAt: pgtype.Timestamp{Time: o.UpdatedAt, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (r *OrderRepository) getOrderForUpdate(ctx context.Context, q *Queries, orderID uuid.UUID) (*order.Order, error) {
	o, err := q.GetOrderForUpdate(ctx, GetOrderForUpdateParams{
		OrderID: orderID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.NewErrNotFound(fmt.Sprintf("order with ID %s", orderID))
		}
		return nil, err
	}

	return order.Load(order.LoadParams{
		OrderID:    o.ID,
		Status:     o.Status,
		CreatedAt:  o.CreatedAt.Time,
		UpdatedAt:  o.UpdatedAt.Time,
		CustomerID: o.CustomerID,
	}), nil
}
