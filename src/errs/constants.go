package errs

const (
	MissedEnvVarErr                = "Error you have to provide proper environment variable: "
	InternalErr                    = "Internal error happened"
	UnprocessableRequestErr        = "Unprocessable request"
	UsernameUniquenessViolationErr = "This username is already in use"
	ShortPasswordErr               = "Password should be at least 8 characters"
	MismatchedPasswords            = "provided password is incorrect"
)
