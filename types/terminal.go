package types

import (
	"fmt"
	"time"
)

type Status string

const (
	StatusOffline  Status = "OFFLINE"
	StatusOnline   Status = "ONLINE"
	StatusLocked   Status = "LOCKED"
	StatusBusy     Status = "BUSY"
	StatusLoggedIn Status = "LOGGED_IN"
)

type Terminal struct {
	model

	Name     string `json:"name" gorm:"primary_key"`
	Hostname string `json:"hostname"`

	RoomID          uint   `json:"room_id"`
	Room            Room   `json:"-"`
	PositionX       uint   `json:"pos_x"`
	PositionY       uint   `json:"pos_y"`
	OperatingSystem string `json:"operating_system"`

	Status Status `json:"status" sql:"-"`
}

const statusKey = "directrd.terminal.%s.status"

/* These are GORM hooks, see here:
   http://gorm.io/docs/hooks.html  */

func (t *Terminal) AfterFind() error {
	if res := ctx.Redis().Get(fmt.Sprintf(statusKey, t.Name)); res.Err() == nil {
		t.Status = Status(res.Val())
	} else {
		t.Status = StatusOffline
	}
	return nil
}

func (t *Terminal) BeforeSave() error {
	ctx.Redis().Set(fmt.Sprintf(statusKey, t.Name), t.Status, time.Hour*1)
	return nil
}
