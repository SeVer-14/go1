package v1

import (
	"github.com/gin-gonic/gin"
	"go1/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	h.initProductRoutes(api)
	h.initCartRoutes(api)
	h.initOrderRoutes(api)
}
