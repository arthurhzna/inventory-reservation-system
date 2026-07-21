package rule

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
)

func Equal(
	field string,
	value string,
	compare string,
) core.Rule {

	return func() core.FieldError {

		if value != compare {

			return core.NewValidationError(
				field,
				constant.TagEqual,
				"",
			)
		}

		return nil
	}
}
