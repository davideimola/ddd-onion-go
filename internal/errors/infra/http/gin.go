package httpInfraErrors

import (
	"davideimola.dev/ddd-onion/internal/errors/app"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func HandleGinHTTPErrors(c *gin.Context, err error) {
	if err == nil {
		return
	}

	if errors.Is(err, apperrors.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	if errors.Is(err, apperrors.ErrInvalidArgument) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	if errors.Is(err, apperrors.ErrAlreadyExists) {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
	}

	if errors.Is(err, apperrors.ErrPreconditionFailed) {
		c.JSON(http.StatusPreconditionFailed, gin.H{"error": err.Error()})
	}

	if errors.Is(err, apperrors.ErrUnauthenticated) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	logrus.WithError(err).Error("unexpected error occurred")
}
