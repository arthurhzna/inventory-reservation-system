package enum

const (
	RoleAdmin    = "ADMIN"
	RoleCustomer = "CUSTOMER"
)

const (
	RoleAdminID    int64 = 1
	RoleCustomerID int64 = 2
)

var RoleNameToID = map[string]int64{
	RoleAdmin:    RoleAdminID,
	RoleCustomer: RoleCustomerID,
}

var RoleIDToName = map[int64]string{
	RoleAdminID:    RoleAdmin,
	RoleCustomerID: RoleCustomer,
}
