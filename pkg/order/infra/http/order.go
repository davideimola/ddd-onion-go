package orderHTTPService

import (
	httpInfraErrors "davideimola.dev/ddd-onion/internal/errors/infra/http"
	orderService "davideimola.dev/ddd-onion/pkg/order/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type OrderHTTPService struct {
	orderService *orderService.OrderService
}

func New(service *orderService.OrderService) *OrderHTTPService {
	return &OrderHTTPService{
		orderService: service,
	}
}

type PostOrderRequest struct {
	CustomerID uuid.UUID `json:"customer_id"`
}

func (s *OrderHTTPService) PostOrder(c *gin.Context) {
	var request PostOrderRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := s.orderService.Create(c.Request.Context(), request.CustomerID)
	if err != nil {
		httpInfraErrors.HandleGinHTTPErrors(c, err)
		return
	}

	c.JSON(http.StatusCreated, order)
}
