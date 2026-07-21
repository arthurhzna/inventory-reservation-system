package validation

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/request"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/rule"
)

func CreateDeviceRules(
	req *request.CreateDeviceRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredString(
			constant.DeviceNameField,
			req.Name,
		),
	}
}
