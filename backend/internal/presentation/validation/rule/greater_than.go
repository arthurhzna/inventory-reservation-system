package rule

import (
	"strconv"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
)

func GreaterThan(
	field string,
	value int64,
	min int64,
) core.Rule {

	return func() core.FieldError {

		if value <= min {

			return core.NewValidationError(
				field,
				constant.TagGreaterThan,
				strconv.FormatInt(value, 10),
			)
		}

		return nil
	}
}
