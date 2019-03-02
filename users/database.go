package users

import (
	"github.com/enderian/directrd/types"
)

func findUser(user *types.User) error {
	return ctx.DB().Model(user).Where("username = ?", user.Username).Select()
}

func updateUser(user *types.User) error {
	_, err := ctx.DB().Model(user).Where("username = ?", user.Username).Update()
	return err
}

func insertUser(user *types.User) error {
	_, err := ctx.DB().Model(user).Insert()
	return err
}
