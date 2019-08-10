package types

import (
	"github.com/jinzhu/gorm"
	"time"
)

type SessionStatus int

const (
	SessionStarting SessionStatus = iota
	SessionStarted
	SessionEnded
)

type Session struct {
	gorm.Model

	InternalId string `json:"internal_id" gorm:"primary_key"`
	SessionId  string `json:"account_id"`

	Machine  string `json:"machine"`
	Username string `json:"username"`
	User     *User  `json:"user" sql:"-"`

	Status  SessionStatus `json:"status,notnull"`
	Expires time.Time     `json:"expires"`
}
