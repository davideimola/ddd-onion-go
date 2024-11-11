package pgInventoryRepository

import (
	"context"
	apperrors "davideimola.dev/ddd-onion/internal/errors/app"
	"davideimola.dev/ddd-onion/pkg/inventory"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"strconv"
)

type ProductRepository struct {
	db *pgxpool.Pool
}

func (p ProductRepository) GetProductByID(ctx context.Context, id uuid.UUID) (*inventory.Product, error) {
	q := New(p.db)

	prod, err := q.GetProductByID(ctx, GetProductByIDParams{
		ProductID: id,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.NewErrNotFound(fmt.Sprintf("product with id %s", id))
		}

		return nil, err
	}

	return toProductModel(prod)
}

func (p ProductRepository) UpdateProduct(ctx context.Context, productID uuid.UUID, updateFn func(p *inventory.Product) error) (*inventory.Product, error) {
	tx, err := p.db.Begin(ctx)
	if err != nil {
		return nil, err
	}

	defer tx.Rollback(ctx)

	q := New(tx)

	prod, err := p.getProductForUpdate(ctx, q, productID)
	if err != nil {
		return nil, err
	}

	err = updateFn(prod)
	if err != nil {
		return nil, err
	}

	price := pgtype.Numeric{}
	parse := strconv.FormatFloat(prod.Price, 'f', -1, 64)
	if err := price.Scan(parse); err != nil {
		return nil, err
	}

	err = q.UpdateProduct(ctx, UpdateProductParams{
		ProductID: prod.ID,
		Name:      prod.Name,
		Price:     price,
		Quantity:  int32(prod.Quantity()),
		UpdatedAt: pgtype.Timestamp{Time: prod.UpdatedAt, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return prod, nil
}

func (p ProductRepository) getProductForUpdate(ctx context.Context, q *Queries, productID uuid.UUID) (*inventory.Product, error) {
	prod, err := q.GetProductForUpdate(ctx, GetProductForUpdateParams{
		ProductID: productID,
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.NewErrNotFound(fmt.Sprintf("product with id %s", productID))
		}

		return nil, err
	}

	return toProductModel(prod)
}

func toProductModel(p Product) (*inventory.Product, error) {
	price, err := p.Price.Float64Value()
	if err != nil {
		return nil, err
	}
	return inventory.LoadProduct(inventory.LoadProductParams{
		ID:        p.ID,
		Name:      p.Name,
		Price:     price.Float64,
		Quantity:  int(p.Quantity),
		CreatedAt: p.CreatedAt.Time,
		UpdatedAt: p.UpdatedAt.Time,
	}), nil
}

func NewProductRepository(db *pgxpool.Pool) inventory.ProductRepository {
	return &ProductRepository{
		db: db,
	}
}
