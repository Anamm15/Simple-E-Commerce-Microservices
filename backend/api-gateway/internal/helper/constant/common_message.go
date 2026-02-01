package constant

const (
	// Success Messages
	MsgSuccess = "Operation successful"
	MsgCreated = "Resource created successfully"
	MsgUpdated = "Resource updated successfully"
	MsgDeleted = "Resource deleted successfully"

	// Client Error Messages
	MsgBadRequest          = "The request could not be understood or contains invalid parameters"
	MsgInvalidRequest      = "Invalid request body or format"
	MsgUnauthorized        = "Authentication is required to access this resource"
	MsgForbidden           = "You do not have permission to perform this action"
	MsgNotFound            = "The requested resource was not found"
	MsgUnprocessableEntity = "The server understands the content type but cannot process the instructions"

	// Server Error Messages
	MsgInternalServerError = "An unexpected error occurred on our end. Please try again later"
	MsgServiceUnavailable  = "The server is currently unable to handle the request"
)
