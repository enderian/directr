package types

import (
	"github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	model

	Username    string `json:"username" sql:",pk"`
	FullName    string `json:"full_name"`
	Affiliation string `json:"affiliation"`

	PhoneNumber string `json:"phone_number"`

	DN string `json:"dn,omitempty"`

	Local    bool   `json:"local,omitempty"`
	Password []byte `json:"password,omitempty"`

	Extras postgres.Hstore `json:"extras,omitempty"`
}
