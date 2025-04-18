package response

const (
	Success = 2000
	Created = 2001

	ErrorBadRequest       = 4000
	ErrorUnauthorized     = 4001
	ErrorForbidden        = 4003
	ErrorNotFound         = 4004
	ErrorParamInvalid     = 4005
	ErrorMethodNotAllowed = 4006
	ErrorInvalidOTP       = 4007
	ErrorSendEmailOTP     = 4008

	ErrorInternalServerError = 5000
	ErrorServiceUnavailable  = 5003
	// Register code
	ErrorCodeUserHasExists = 5004
)

var msg = map[int]string{
	Success: "Success",
	Created: "Created",

	ErrorBadRequest:       "Bad Request",
	ErrorUnauthorized:     "Unauthorized",
	ErrorForbidden:        "Forbidden",
	ErrorNotFound:         "Not Found",
	ErrorParamInvalid:     "Param Invalid",
	ErrorMethodNotAllowed: "Method Not Allowed",
	ErrorInvalidOTP:       "Invalid OTP",
	ErrorSendEmailOTP:     "Error send Email OTP",

	ErrorInternalServerError: "Internal Server Error",
	ErrorServiceUnavailable:  "Service Unavailable",

	ErrorCodeUserHasExists: "User has already exists",
}
