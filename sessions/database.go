package sessions

import (
	"errors"
	"github.com/enderian.directrd/types"
)

func findSession(id *types.Session) error {
	return errors.New("session not found")
}

func findSessions(id *types.Session) ([]*types.Session, error) {
	return nil, errors.New("session not found")
}

func insertSession(session *types.Session) error {
	return nil
}

func updateSession(session *types.Session) error {
	return errors.New("session not found")
}
