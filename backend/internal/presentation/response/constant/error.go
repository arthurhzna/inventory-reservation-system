package constant

const (
	InternalServerErrorMessage     = "currently our server is facing unexpected error, please try again later"
	EOFErrorMessage                = "missing body request"
	JsonSyntaxErrorMessage         = "invalid JSON syntax"
	JsonUnmarshallTypeErrorMessage = "invalid value for %s"
	UnauthorizedErrorMessage       = "unauthorized"
	RequestTimeoutErrorMessage     = "failed to process request in time, please try again"
	ForbiddenErrorMessage          = "you are not allowed to access this resource"
	ValidationErrorMessage         = "input validation error"
	NotFoundErrorMessage           = "%s not found"
	ConflictErrorMessage           = "%s already exists"
	BadRequestErrorMessage         = "%s is invalid"
)
