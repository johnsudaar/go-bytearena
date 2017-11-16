package models

const (
	HandshakeType = "Handshake"
	MutationType  = "Mutation"
)

// TODO: Find a better name for that
type Method struct {
	AgentID string      `json:"agentid"`
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}
