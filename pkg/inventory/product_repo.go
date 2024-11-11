package inventory

import (
	"context"
	"github.com/google/uuid"
)

type ProductRepository interface {
	GetProductByID(ctx context.Context, id uuid.UUID) (*Product, error)
	UpdateProduct(ctx context.Context, productID uuid.UUID, updateFn func(p *Product) error) (*Product, error)
}
