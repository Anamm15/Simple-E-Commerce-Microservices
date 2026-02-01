package enum

type AccountRole string

const (
	UserRoleSuperAdmin AccountRole = "super_admin"
	UserRoleAdmin      AccountRole = "admin"
	UserRoleCustomer   AccountRole = "user"
)
