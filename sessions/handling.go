package sessions

import (
	"github.com/Knetic/govaluate"
	"github.com/enderian/directrd/types"
	"log"
	"time"
)

func Authorize(user *types.User, machine string) error {
	if conf == nil {
		return nil
	}

	//If there are authorization rules.
	if conf.User.AuthorizationRules != nil {
		concurrent := 0

		err := conf.User.AuthorizationRules.ExecuteRules(govaluate.MapParameters{
			"user":       user,
			"machine":    machine,
			"concurrent": concurrent,
		})
		if err != nil {
			return err
		}
	}

	if err := insertSession(&types.Session{
		Username: user.Username,
		Machine:  machine,
		Expires:  time.Now().Add(time.Minute * 5),
	}); err != nil {
		log.Printf("Error while creating session: %s", err.Error())
		return err
	}

	//Successful session.
	return nil
}

func Start(username string, machine string) error {
	if conf == nil {
		return nil
	}
	session := &types.Session{Machine: machine}
	if err := findSession(session); err == nil && session.Username == username {
		session.Status = types.SessionStarted
		if err := updateSession(session); err != nil {
			log.Printf("Error while updating session whilst starting: %s", err.Error())
			return err
		}
	} else {
		if err := insertSession(&types.Session{
			Username: username,
			Machine:  machine,
			Expires:  time.Now().Add(time.Minute * 5),
		}); err != nil {
			log.Printf("Error while creating session whilst starting: %s", err.Error())
			return err
		}
	}
	return nil
}

func Update(username string, machine string) error {
	if conf == nil {
		return nil
	}
	session := &types.Session{Machine: machine}
	if err := findSession(session); err == nil && session.Username == username {
		session.Expires = time.Now().Add(time.Minute * 3)
		if err := updateSession(session); err != nil {
			log.Printf("Error while updating session whilst updating: %s", err.Error())
			return err
		}
	} else {
		if err := insertSession(&types.Session{
			Username: username,
			Machine:  machine,
			Status:   types.SessionStarted,
			Expires:  time.Now().Add(time.Minute * 3),
		}); err != nil {
			log.Printf("Error while creating session whilst updating: %s", err.Error())
			return err
		}
	}
	return nil
}

func End(username string, machine string) error {
	if conf == nil {
		return nil
	}
	session := &types.Session{Machine: machine}
	if err := findSession(session); err == nil {
		session.Status = types.SessionEnded
		if err := updateSession(session); err != nil {
			log.Printf("Error while updating session whilst completing: %s", err.Error())
			return err
		}
	}
	return nil
}
