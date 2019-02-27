package types

type dirError string

func (d dirError) Error() string {
	return string(d)
}

const (
	ErrorConfiguration dirError = "There is an error in the server configuration. Contact an administrator."
	ErrorCredentials   dirError = "Invalid credentials."
	ErrorNotRegistered dirError = "You are not registered. Please contact an administrator if you think this is a mistake."
	ErrorNotAuthorized dirError = "You are not authorized to access this device. Please contact an administrator if you think this is a mistake."
)
