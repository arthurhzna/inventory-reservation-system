package rule

import (
	"strconv"
	"strings"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
)

func RequiredString(
	field string,
	value string,
) core.Rule {

	return func() core.FieldError {

		if strings.TrimSpace(value) == "" {

			return core.NewValidationError(
				field,
				constant.TagRequired,
				value,
			)
		}

		return nil
	}
}

func RequiredInt64(
	field string,
	value int64,
) core.Rule {
	return func() core.FieldError {

		if value == 0 {

			return core.NewValidationError(
				field,
				constant.TagRequired,
				strconv.FormatInt(value, 10),
			)
		}

		return nil
	}
}
