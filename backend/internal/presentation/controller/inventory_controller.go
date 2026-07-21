package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/request"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/usecase"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/response"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation"
)

type InventoryController struct {
	inventoryUseCase usecase.InventoryUseCaseInterface
}

func NewInventoryController(
	inventoryUseCase usecase.InventoryUseCaseInterface,
) *InventoryController {
	return &InventoryController{
		inventoryUseCase: inventoryUseCase,
	}
}

func (c *InventoryController) Reserve(ctx *gin.Context) {

	req := new(request.ReserveInventoryRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	if err := validation.Validate(
		validation.ReserveInventoryRules(req),
	); err != nil {
		ctx.Error(err)
		return
	}

	res, err := c.inventoryUseCase.Reserve(
		ctx.Request.Context(),
		req,
	)

	if err != nil {
		ctx.Error(response.MapError(err))
		return
	}

	response.ResponseCreated(
		ctx,
		res,
	)
}

func (c *InventoryController) Confirm(ctx *gin.Context) {

	req := new(request.ConfirmReservationRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	if err := validation.Validate(
		validation.ConfirmReservationRules(req),
	); err != nil {
		ctx.Error(err)
		return
	}

	res, err := c.inventoryUseCase.Confirm(
		ctx.Request.Context(),
		req,
	)

	if err != nil {
		ctx.Error(response.MapError(err))
		return
	}

	response.ResponseOK(
		ctx,
		res,
	)
}

func (c *InventoryController) GetStock(ctx *gin.Context) {

	itemID := ctx.Query("item_id")

	res, err := c.inventoryUseCase.GetStock(
		ctx.Request.Context(),
		itemID,
	)

	if err != nil {
		ctx.Error(response.MapError(err))
		return
	}

	response.ResponseOK(
		ctx,
		res,
	)
}

func (c *InventoryController) ListInventory(ctx *gin.Context) {

	res, err := c.inventoryUseCase.ListInventory(
		ctx.Request.Context(),
	)

	if err != nil {
		ctx.Error(response.MapError(err))
		return
	}

	response.ResponseOK(
		ctx,
		res,
	)
}
