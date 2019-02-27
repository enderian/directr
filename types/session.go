package types

import (
	"time"
)

type SessionStatus int

const (
	SessionStarting SessionStatus = iota
	SessionStarted
	SessionEnded
)

type Session struct {
	Machine string `json:"machine"`

	Username string `json:"username"`
	User     *User  `json:"user"`

	Status  SessionStatus `json:"status"`
	Expires time.Time     `json:"expires"`
}
