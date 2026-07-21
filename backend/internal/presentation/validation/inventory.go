package validation

import (
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/application/dto/request"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/domain/policy"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/constant"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/core"
	"github.com/arthurhzna/inventory-reservation-system/backend/internal/presentation/validation/rule"
)

func ReserveInventoryRules(
	req *request.ReserveInventoryRequest,
) []core.Rule {

	return []core.Rule{

		rule.RequiredString(
			constant.InventoryUserIDField,
			req.UserID,
		),

		rule.RequiredString(
			constant.InventoryItemIDField,
			req.ItemID,
		),

		rule.RequiredInt64(
			constant.InventoryQuantityField,
			req.Quantity,
		),

		rule.GreaterThan(
			constant.InventoryQuantityField,
			req.Quantity,
			policy.MinQuantity,
		),
	}
}
