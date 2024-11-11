package inventory

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID        uuid.UUID
	Name      string
	Price     float64
	quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LoadProductParams struct {
	ID        uuid.UUID
	Name      string
	Price     float64
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func LoadProduct(params LoadProductParams) *Product {
	return &Product{
		ID:        params.ID,
		Name:      params.Name,
		Price:     params.Price,
		quantity:  params.Quantity,
		CreatedAt: params.CreatedAt,
		UpdatedAt: params.UpdatedAt,
	}
}

func (p *Product) Quantity() int {
	return p.quantity
}

func (p *Product) AddQuantity(q int) {
	p.quantity += q
}

func (p *Product) SellQuantity(q int) {
	p.quantity -= q
}
