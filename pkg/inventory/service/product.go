package inventoryService

import (
	"context"
	"davideimola.dev/ddd-onion/pkg/inventory"
	"github.com/google/uuid"
)

type ProductService struct {
	productRepo inventory.ProductRepository
}

func New(productRepo inventory.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

func (s *ProductService) GetProductByID(ctx context.Context, id uuid.UUID) (*inventory.Product, error) {
	return s.productRepo.GetProductByID(ctx, id)
}

func (s *ProductService) AddProductQuantity(ctx context.Context, productID uuid.UUID, quantity int) (*inventory.Product, error) {
	return s.productRepo.UpdateProduct(ctx, productID, func(p *inventory.Product) error {
		p.AddQuantity(quantity)
		return nil
	})
}

func (s *ProductService) SellProductQuantity(ctx context.Context, productID uuid.UUID, quantity int) (*inventory.Product, error) {
	return s.productRepo.UpdateProduct(ctx, productID, func(p *inventory.Product) error {
		p.SellQuantity(quantity)
		return nil
	})
}
