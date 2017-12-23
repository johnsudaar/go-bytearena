package models

const (
	HandshakeType = "Handshake"
	ActionsType   = "Actions"
)

// TODO: Find a better name for that
type Method struct {
	AgentID string      `json:"agentid"`
	Method  string      `json:"method"`
	Payload interface{} `json:"payload"`
}
