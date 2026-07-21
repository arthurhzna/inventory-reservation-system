package validation

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/request"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/rule"
)

func ConfirmReservationRules(
	req *request.ConfirmReservationRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredString(
			constant.InventoryReservationIDField,
			req.ReservationID,
		),

		rule.UUID(
			constant.InventoryReservationIDField,
			req.ReservationID,
		),
	}
}
