package utils

import (
	"fmt"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
)

func TagToMsg(fe core.FieldError) string {

	switch fe.Tag() {

	case constant.TagRequired:
		return fmt.Sprintf(
			"%s is required",
			fe.Field(),
		)

	case constant.TagEmail:
		return fmt.Sprintf(
			"%s has invalid email format",
			fe.Field(),
		)

	case constant.TagMinLength:
		return fmt.Sprintf(
			"%s length must be at least %s",
			fe.Field(),
			fe.Param(),
		)

	case constant.TagEqual:
		return fmt.Sprintf(
			"%s is invalid",
			fe.Field(),
		)

	case constant.TagTimeFormat:
		return fmt.Sprintf(
			"%s has invalid time format",
			fe.Field(),
		)

	default:
		return "invalid input"
	}
}
