package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/request"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/usecase"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/response"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation"
)

type UserController struct {
	userUseCase usecase.UserUseCaseInterface
}

func NewUserController(
	userUseCase usecase.UserUseCaseInterface,
) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (c *UserController) Register(ctx *gin.Context) {

	req := new(request.RegisterUserRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	if err := validation.Validate(
		validation.RegisterUserRules(req),
	); err != nil {
		ctx.Error(err)
		return
	}

	res, err := c.userUseCase.Register(
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

func (c *UserController) Login(ctx *gin.Context) {

	req := new(request.LoginUserRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	if err := validation.Validate(
		validation.LoginUserRules(req),
	); err != nil {
		ctx.Error(err)
		return
	}

	res, err := c.userUseCase.Login(
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
