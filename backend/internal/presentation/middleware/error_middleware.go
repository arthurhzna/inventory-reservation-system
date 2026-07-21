package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/response"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/response/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/response/dto"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/utils"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		errLen := len(ctx.Errors)
		if errLen > 0 {
			err := ctx.Errors.Last()

			switch e := err.Err.(type) {
			case validation.Errors:
				handleValidationError(ctx, e)
			case *json.SyntaxError:
				handleJsonSyntaxError(ctx)
			case *json.UnmarshalTypeError:
				handleJsonUnmarshalTypeError(ctx, e)
			case *time.ParseError:
				handleParseTimeError(ctx, e)
			case *response.ResponseError:
				ctx.AbortWithStatusJSON(e.GetCode(), dto.WebResponse[any]{
					Message: e.DisplayMessage(),
				})
			default:
				if errors.Is(e, io.EOF) {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.WebResponse[any]{
						Message: constant.EOFErrorMessage,
					})
					return
				}

				ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.WebResponse[any]{
					Message: constant.InternalServerErrorMessage,
				})
			}
		}
	}
}

func handleJsonSyntaxError(ctx *gin.Context) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: constant.JsonSyntaxErrorMessage,
	})
}

func handleJsonUnmarshalTypeError(ctx *gin.Context, err *json.UnmarshalTypeError) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: fmt.Sprintf(constant.JsonUnmarshallTypeErrorMessage, err.Field),
	})
}

func handleParseTimeError(ctx *gin.Context, err *time.ParseError) {
	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: fmt.Sprintf("please send time in format of %s, got: %s", utils.ConvertGoTimeLayoutToReadable(err.Layout), err.Value),
	})
}

func handleValidationError(ctx *gin.Context, err validation.Errors) {
	ve := []dto.FieldError{}

	for _, fe := range err {
		ve = append(ve, dto.FieldError{
			Field:   fe.Field(),
			Message: utils.TagToMsg(fe),
		})
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.WebResponse[any]{
		Message: constant.ValidationErrorMessage,
		Errors:  ve,
	})
}
