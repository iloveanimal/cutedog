package errorHandle

type GeneralErrorInter interface {
	Error() string
	Location() string
}

type GeneralError struct {
	location string // Location where the error occurs
	message  string // Error Message
}

func GetGeneralError(location, message string) GeneralError {
	return GeneralError{
		location: location,
		message:  message,
	}
}

func (g GeneralError) Error() string {
	return g.message
}

func (g GeneralError) Location() string {
	return g.location
}
