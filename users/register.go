package users

import (
	"github.com/Knetic/govaluate"
	"github.com/enderian/directrd/types"
)

func autoRegister(user *types.User, identifier string) (err error) {
	if !ctx.Conf().User.AutoRegister {
		return types.ErrorNotRegistered
	}

	//defers run at the end of the function. Like finally!
	defer func() {
		if err == nil {
			//Insert the user if all else went well!
			err = insertUser(user)
		}
	}()
	if ctx.Conf().User.AutoRegisterRules == nil {
		return nil
	}
	//Check ALL THE RULES.
	return ctx.Conf().User.AutoRegisterRules.ExecuteRules(govaluate.MapParameters{
		"user":    user,
		"machine": identifier,
	})
}
