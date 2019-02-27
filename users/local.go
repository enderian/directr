package directoryUsers

import (
	"ender.gr/directrd/types"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

/**
authenticateLocal Authenticates a user using their stored password.
If local is true, it also requires the user was created locally.
*/
func authenticateLocal(user types.User, password string, local bool) error {
	if err := findUser(&user); err != nil {
		log.Printf("Error while authenticating local %s: %s", user.Username, err.Error())
		return err
	}
	if local && !user.Local {
		return errors.New("user is not a local user")
	}
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
