package types

type User struct {
	Username    string `json:"username"`
	FullName    string `json:"full_name"`
	Affiliation string `json:"affiliation"`

	DN string `json:"dn"`

	Local    bool   `json:"local"`
	Password []byte `json:"password"`

	Extras map[string]string `json:"extras"`
}
