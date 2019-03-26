package types

type CommandType int

const (
	CommandShutdown CommandType = iota + 1
	CommandRestart
	CommandLogout // [Username]
)

type Command struct {
	Type      CommandType `json:"type"`
	Arguments []string    `json:"arguments"`
}
