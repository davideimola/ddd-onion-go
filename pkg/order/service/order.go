package orderService

import (
	"context"
	apperrors "davideimola.dev/ddd-onion/internal/errors/app"
	"davideimola.dev/ddd-onion/pkg/order"
	"github.com/google/uuid"
)

type OrderService struct {
	orderRepository order.OrderRepository
	productService  ProductService
}

func New(orderRepository order.OrderRepository, productService ProductService) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
		productService:  productService,
	}
}

func (s *OrderService) Create(ctx context.Context, customerID uuid.UUID, items map[uuid.UUID]int) (*order.Order, error) {
	o := order.New(customerID, items)

	for itemID, quantity := range items {
		prod, err := s.productService.GetProductByID(ctx, itemID)
		if err != nil {
			return nil, err
		}

		if prod.Quantity() < quantity {
			return nil, apperrors.NewErrPreconditionFailed("not enough quantity in stock")
		}

		// NOTE: Of course, this is not a real-world scenario. In a real-world scenario, we would work with a transaction
		_, err = s.productService.SellProductQuantity(ctx, itemID, quantity)
		if err != nil {
			return nil, err
		}
	}

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
		for itemID, quantity := range o.Items() {
			_, err := s.productService.AddProductQuantity(ctx, itemID, quantity)
			if err != nil {
				return err
			}
		}
		return o.Cancel()
	})
	if err != nil {
		return nil, err
	}

	return o, nil
}
