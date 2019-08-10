package users

import (
	"errors"
	"github.com/enderian/directrd/types"
	"golang.org/x/crypto/bcrypt"
	"log"
)

/**
authenticateLocal Authenticates a user using their stored password.
If local is true, it also requires the user was created locally.
*/
func authenticateLocal(user types.User, password string, local bool) error {
	if err := ctx.DB().Find(&user).Error; err != nil {
		log.Printf("Error while authenticating local %s: %s", user.Username, err)
		return err
	}
	if local && !user.Local {
		return errors.New("user is not a local user")
	}
	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
