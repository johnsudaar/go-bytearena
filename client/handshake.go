package client

import (
	"github.com/johnsudaar/go-bytearena/models"
	"github.com/pkg/errors"
)

func (c *ChanClient) Handshake() error {
	err := c.Send(models.HandshakeType, models.Handshake{
		Version: c.Version,
	})
	if err != nil {
		return errors.Wrap(err, "fail to start handshake")
	}
	return nil
}
