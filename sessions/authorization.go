package sessions

import (
	"github.com/Knetic/govaluate"
	"github.com/enderian/directrd/types"
	"github.com/iris-contrib/go.uuid"
	"log"
	"time"
)

func Authorize(user *types.User, machine string) error {
	if ctx.Conf() == nil {
		return nil
	}

	//If there are authorization rules.
	if ctx.Conf().User.AuthorizationRules != nil {
		sessions, err := findSessions(&types.Session{Username: user.Username})
		if err != nil {
			log.Printf("Error while retrieving sessions from database: %s", err.Error())
		}
		err = ctx.Conf().User.AuthorizationRules.ExecuteRules(govaluate.MapParameters{
			"user":       user,
			"machine":    machine,
			"concurrent": len(sessions),
		})
		if err != nil {
			return err
		}
	}

	if err := ctx.DB().Create(&types.Session{
		InternalId: uuid.Must(uuid.NewV4()).String(),
		Username:   user.Username,
		Machine:    machine,
		Expires:    time.Now().Add(time.Minute * 5),
	}).Error; err != nil {
		log.Printf("Error while creating session: %s", err)
		return err
	}

	//Successful session.
	return nil
}

func Start(sessionId string, username string, machine string) error {
	if ctx.Conf() == nil {
		return nil
	}
	session := &types.Session{Machine: machine}
	if err := findSession(session); err == nil && session.Username == username {
		session.SessionId = sessionId
		session.Status = types.SessionStarted
		if err := ctx.DB().Save(session).Error; err != nil {
			log.Printf("Error while updating session whilst starting: %s", err.Error())
			return err
		}
	} else {
		if err := ctx.DB().Create(&types.Session{
			InternalId: uuid.Must(uuid.NewV4()).String(),
			SessionId:  sessionId,
			Username:   username,
			Machine:    machine,
			Expires:    time.Now().Add(time.Minute * 5),
		}).Error; err != nil {
			log.Printf("Error while creating session whilst starting: %s", err)
			return err
		}
	}
	log.Printf("Session for user %s on %s started.", username, machine)
	return nil
}

func Update(sessionId string, username string, machine string) error {
	if ctx.Conf() == nil {
		return nil
	}
	session := &types.Session{SessionId: sessionId}
	if err := findSession(session); err == nil && session.Username == username {
		session.Expires = time.Now().Add(time.Minute * 3)
		if err := ctx.DB().Save(session).Error; err != nil {
			log.Printf("Error while updating session whilst updating: %s", err)
			return err
		}
	} else {
		if err := ctx.DB().Create(&types.Session{
			InternalId: uuid.Must(uuid.NewV4()).String(),
			SessionId:  sessionId,
			Username:   username,
			Machine:    machine,
			Status:     types.SessionStarted,
			Expires:    time.Now().Add(time.Minute * 3),
		}).Error; err != nil {
			log.Printf("Error while creating session whilst updating: %s", err)
			return err
		}
	}
	log.Printf("Session for user %s on %s updated.", username, machine)
	return nil
}

func End(sessionId string, username string, machine string) error {
	if ctx.Conf() == nil {
		return nil
	}
	session := &types.Session{SessionId: sessionId}
	if err := findSession(session); err == nil {
		session.Status = types.SessionEnded
		if err := ctx.DB().Save(session).Error; err != nil {
			log.Printf("Error while updating session whilst completing: %s", err)
			return err
		}
	}
	log.Printf("Session for user %s on %s ended.", username, machine)
	return nil
}
