package sessions

import (
	"github.com/enderian/directrd/types"
)

func findSession(session *types.Session) error {
	return ctx.DB().Where("internal_id = ?", session.InternalId).Find(&session).Error
}

func findSessions(session *types.Session) ([]*types.Session, error) {
	var sessions []*types.Session
	err := ctx.DB().
		Where("machine = ? OR username = ?", session.Machine, session.Username).
		Find(&sessions).
		Error
	return sessions, err
}