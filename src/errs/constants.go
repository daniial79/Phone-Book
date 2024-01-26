package errs

const (
	MissedEnvVarErr                = "Error you have to provide proper environment variable: "
	InternalErr                    = "Internal error happened"
	UnprocessableRequestErr        = "Unprocessable request"
	UsernameUniquenessViolationErr = "This username is already in use"
	ShortPasswordErr               = "Password should be at least 8 characters"
	InsufficientCredentialsErr     = "You have to provide enough credentials"
	BadRequestErr                  = "Bad request"
	InvalidRefreshTokenErr         = "Invalid refresh token"
	UnauthorizedErr                = "Unauthorized user"
	CookieNotFoundErr              = "Cookie not found"
	UserNotFoundErr                = "There is no user with such username"
)
