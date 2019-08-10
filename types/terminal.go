package types

import (
	"github.com/jinzhu/gorm"
)



type Terminal struct {
	gorm.Model

	Name     string `json:"name" gorm:"primary_key"`
	Hostname string `json:"hostname"`

	OperatingSystem string `json:"operating_system"`
}
