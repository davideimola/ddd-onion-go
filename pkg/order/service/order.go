package orderService

import (
	"context"
	"davideimola.dev/ddd-onion/pkg/order"
	"github.com/google/uuid"
)

type OrderService struct {
	orderRepository order.OrderRepository
}

func New(orderRepository order.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (s *OrderService) Create(ctx context.Context, customerID uuid.UUID) (*order.Order, error) {
	o := order.New(customerID)

	err := s.orderRepository.CreateOrder(ctx, o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (s *OrderService) Ship(ctx context.Context, orderID uuid.UUID) (*order.Order, error) {
	o, err := s.orderRepository.UpdateOrder(ctx, orderID, func(o *order.Order) error {
		return o.Ship()
	})
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (s *OrderService) Cancel(ctx context.Context, orderID uuid.UUID) (*order.Order, error) {
	o, err := s.orderRepository.UpdateOrder(ctx, orderID, func(o *order.Order) error {
		return o.Cancel()
	})
	if err != nil {
		return nil, err
	}

	return o, nil
}
