package rule

import (
	"strings"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
)

func OneOf(
	field string,
	value string,
	allowed []string,
) core.Rule {

	return func() core.FieldError {

		if strings.TrimSpace(value) == "" {
			return nil
		}

		for _, v := range allowed {
			if value == v {
				return nil
			}
		}

		return core.NewValidationError(
			field,
			constant.TagOneOf,
			strings.Join(allowed, " "),
		)
	}
}
