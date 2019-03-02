package sessions

import (
	"github.com/enderian/directrd/types"
)

func findSession(session *types.Session) error {
	return ctx.DB().Model(session).Where("internal_id = ?", session.InternalId).Select()
}

func findSessions(session *types.Session) ([]*types.Session, error) {
	var sessions []*types.Session
	err := ctx.DB().Model(&sessions).
		Where("machine = ?", session.Machine).
		WhereOr("username = ?", session.Username).
		Select()
	return sessions, err
}

func insertSession(session *types.Session) error {
	_, err := ctx.DB().Model(session).Insert()
	return err
}

func updateSession(session *types.Session) error {
	_, err := ctx.DB().Model(session).Where("internal_id = ?", session.InternalId).UpdateNotNull()
	return err
}
