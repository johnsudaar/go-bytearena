package agent

import (
	"github.com/johnsudaar/go-bytearena/models"
	"github.com/pkg/errors"
)

func (a *Agent) Handshake() error {
	err := a.Send(models.HandshakeType, models.Handshake{
		Version: a.Version,
	})
	if err != nil {
		return errors.Wrap(err, "fail to start handshake")
	}
	return nil
}
