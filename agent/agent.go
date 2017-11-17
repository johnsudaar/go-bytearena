package agent

import "github.com/johnsudaar/go-bytearena/models"

type Agent interface {
	Tick(perception models.Perception, skipped int) *models.Actions // Called every tick
	Raw(raw []byte, skipped int)                                    // Called when a message is received by the server
	Error(err error, skipped int)                                   // Called when an error is received
	Worker()                                                        // Started when the handshake process is finished
}
