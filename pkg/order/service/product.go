package orderService

import (
	"context"
	"davideimola.dev/ddd-onion/pkg/inventory"
	"github.com/google/uuid"
)

type ProductService interface {
	GetProductByID(ctx context.Context, id uuid.UUID) (*inventory.Product, error)
	AddProductQuantity(ctx context.Context, productID uuid.UUID, quantity int) (*inventory.Product, error)
	SellProductQuantity(ctx context.Context, productID uuid.UUID, quantity int) (*inventory.Product, error)
}
