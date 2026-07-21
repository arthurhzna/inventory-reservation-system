package rule

import (
	"strings"

	"github.com/google/uuid"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
)

func UUID(
	field string,
	value string,
) core.Rule {

	return func() core.FieldError {

		if strings.TrimSpace(value) == "" {
			return nil
		}

		_, err := uuid.Parse(value)
		if err != nil {

			return core.NewValidationError(
				field,
				constant.TagUUID,
				"",
			)
		}

		return nil
	}
}
