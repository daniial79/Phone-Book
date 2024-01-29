package errs

const (
	MissedEnvVarErr                = "Error you have to provide proper environment variable: "
	InternalErr                    = "Internal error happened"
	UnprocessableRequestErr        = "Unprocessable request"
	UsernameUniquenessViolationErr = "This username is already in use"
	ShortPasswordErr               = "Password should be at least 8 characters"
	InsufficientCredentialsErr     = "You have to provide enough credentials"
	BadRequestErr                  = "Bad request"
	InvalidToken                   = "Invalid token"
	UnauthorizedErr                = "Unauthorized user"
	CookieNotFoundErr              = "Cookie not found"
	UserNotFoundErr                = "There is no user with such username"
	ContactNotFoundErr             = "Contact with such id is not found"
	EmailNotFoundErr               = "email not found"
	NumberNotFoundErr              = "phone number not found"
	InvalidIdErr                   = "provided Id in url parameter is invalid"
)
