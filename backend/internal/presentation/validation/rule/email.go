package rule

import (
	"regexp"
	"strings"

	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
)

var emailRegex = regexp.MustCompile(
	`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`,
)

func Email(
	field string,
	value string,
) core.Rule {

	return func() core.FieldError {

		if strings.TrimSpace(value) == "" {
			return nil
		}

		if !emailRegex.MatchString(value) {

			return core.NewValidationError(
				field,
				constant.TagEmail,
				"",
			)
		}

		return nil
	}
}
