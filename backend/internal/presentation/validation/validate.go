package validation

import "github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"

type Errors = core.Errors
type FieldError = core.FieldError
type Rule = core.Rule

func Validate(
	rules []Rule,
) error {

	return core.Execute(rules)
}
