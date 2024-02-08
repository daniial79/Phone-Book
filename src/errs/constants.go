package errs

const (
	ErrMissedEnvVar                = "Error you have to provide proper environment variable: "
	ErrInternal                    = "Internal error happened"
	ErrUnprocessableRequest        = "Unprocessable request"
	ErrUsernameUniquenessViolation = "This username is already in use"
	ErrShortPassword               = "Password should be at least 8 characters"
	ErrInsufficientCredentials     = "You have to provide enough credentials"
	ErrBadRequest                  = "Bad request"
	ErrUnauthorized                = "Unauthorized user"
	ErrCookieNotFound              = "Cookie not found"
	ErrUserNotFound                = "There is no user with such username"
	ErrContactNotFound             = "Contact with such id is not found"
	ErrEmailNotFound               = "email not found"
	ErrNumberNotFound              = "phone number not found"
	ErrUserRole                    = "user role must be 'admin' or 'user'"
	ErrInvalidToken                = "Bearer Token is invalid"
	ErrSigningMethodMismatch       = "signing algorithm mismatch"
)
