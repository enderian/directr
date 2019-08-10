package types

import (
	"github.com/jinzhu/gorm"
)

type Room struct {
	gorm.Model

	Name string `json:"name"`
}