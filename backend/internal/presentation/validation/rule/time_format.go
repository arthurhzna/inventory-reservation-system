package rule

import (
	"time"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
)

func TimeFormat(
	field string,
	value string,
	layout string,
) core.Rule {

	return func() core.FieldError {

		if value == "" {
			return nil
		}

		_, err := time.Parse(layout, value)

		if err != nil {

			return core.NewValidationError(
				field,
				constant.TagTimeFormat,
				"",
			)
		}

		return nil
	}
}
