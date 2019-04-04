package types

type Terminal struct {
	Name     string `json:"name"`
	Hostname string `json:"hostname"`

	OperatingSystem string `json:"operating_system"`
}
