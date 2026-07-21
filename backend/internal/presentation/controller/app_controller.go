package controller

import (
	"net/http"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/response/dto"
	"github.com/gin-gonic/gin"
)

type AppController struct {
}

func NewAppController() *AppController {
	return &AppController{}
}

func (c *AppController) Health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, dto.WebResponse[any]{
		Message: "status UP",
	})
}

func (c *AppController) RouteNotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, dto.WebResponse[any]{
		Message: "route not found",
	})
}

func (c *AppController) MethodNotAllowed(ctx *gin.Context) {
	ctx.JSON(http.StatusMethodNotAllowed, dto.WebResponse[any]{
		Message: "method not allowed",
	})
}
