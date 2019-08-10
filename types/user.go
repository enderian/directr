package types

import (
	"github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	model

	Username    string `json:"username" sql:",pk"`
	FullName    string `json:"full_name"`
	Affiliation string `json:"affiliation"`

	DN string `json:"dn"`

	Local    bool   `json:"local"`
	Password []byte `json:"password"`

	Extras postgres.Hstore `json:"extras"`
}
