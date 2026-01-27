package constants

const (
	// User Success Messages
	MsgUserFound         = "User profile retrieved successfully"
	MsgUserUpdated       = "User profile updated successfully"
	MsgUserPasswordReset = "Password has been reset successfully"

	// User Success Messages
	MsgUserCreated = "User account created successfully"
	MsgUserDeleted = "User account deleted successfully"

	//  User Address Success Messages
	MsgAddressAdded    = "Address added successfully"
	MsgAddressUpdated  = "Address updated successfully"
	MsgAddressDeleted  = "Address deleted successfully"
	MsgAddressNotFound = "Address not found"

	// User Error Messages
	MsgUserNotFound      = "User not found"
	MsgUserAlreadyExists = "User with this email or username already exists"
	MsgUserIncomplete    = "Please complete your profile information"
	MsgInvalidUserId     = "The provided User ID is invalid"

	// User Address Error Messages
	MsgAddressConflict = "Address already exists"
)
