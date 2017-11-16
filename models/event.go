package models

import (
	"encoding/json"
)

const (
	MethodPerception = "perception"
	MethodWelcome    = "welcome"
)

type Event struct {
	Method  string          `json:"method"`
	Payload json.RawMessage `json:"payload"`
}

type ErrorEvent struct {
	Error error
}

type RawEvent struct {
	Value []byte
}
