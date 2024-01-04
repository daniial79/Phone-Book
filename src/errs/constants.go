package errs

const (
	InternalErr                    = "Internal error happened"
	UnprocessableRequestErr        = "Unprocessable request"
	UsernameUniquenessViolationErr = "This username is already in use"
	ShortPasswordErr               = "Password should be at least 8 characters"
	MissedEnvVarErr                = "Error you have to provide proper environment variable: "
)
