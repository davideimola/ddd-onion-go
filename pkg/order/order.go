package order

import (
	"davideimola.dev/ddd-onion/internal/errors/app"
	"github.com/google/uuid"
	"time"
)

type Status string

const (
	StatusShipped  Status = "shipped"
	StatusCanceled Status = "canceled"
	StatusCreated  Status = "created"
)

type Order struct {
	ID         uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CustomerID uuid.UUID
	status     Status
}

func New(customerID uuid.UUID) *Order {
	return &Order{
		ID:         uuid.New(),
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		CustomerID: customerID,
		status:     StatusCreated,
	}
}

type LoadParams struct {
	OrderID    uuid.UUID
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CustomerID uuid.UUID
	Status     Status
}

func Load(params LoadParams) *Order {
	return &Order{
		ID:         params.OrderID,
		CreatedAt:  params.CreatedAt,
		UpdatedAt:  params.UpdatedAt,
		CustomerID: params.CustomerID,
		status:     params.Status,
	}
}

func (o *Order) Status() Status {
	return o.status
}

func (o *Order) Ship() error {
	if o.status != StatusCreated {
		return apperrors.NewErrPreconditionFailed("can't ship the order due to an invalid status")
	}

	o.status = StatusShipped
	o.UpdatedAt = time.Now()

	return nil
}

func (o *Order) Cancel() error {
	if o.status == StatusCanceled {
		return apperrors.NewErrPreconditionFailed("can't cancel an already canceled order")
	}

	if o.status == StatusShipped {
		return apperrors.NewErrPreconditionFailed("can't cancel an already shipped order")
	}

	o.status = StatusCanceled
	o.UpdatedAt = time.Now()

	return nil
}
