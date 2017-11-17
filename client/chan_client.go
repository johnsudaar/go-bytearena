package client

import (
	"os"

	"github.com/johnsudaar/go-bytearena/models"
	"github.com/johnsudaar/go-bytearena/transport"
	"github.com/pkg/errors"
)

const (
	VersionClearBeta = "clear_beta"
)

type ChanClient struct {
	Transport transport.Transport
	ID        string
	Version   string
}

func NewChanClient(version string) (*ChanClient, error) {
	if os.Getenv("AGENTID") == "" {
		return nil, errors.New("no agent id specified")
	}

	t, err := transport.FromEnv()
	if err != nil {
		return nil, err
	}

	c := &ChanClient{
		Transport: t,
		ID:        os.Getenv("AGENTID"),
		Version:   version,
	}

	return c, nil
}

func (c *ChanClient) Start() (chan interface{}, error) {
	serverChan, err := c.Transport.Connect()
	if err != nil {
		return nil, err
	}

	err = c.Handshake()
	if err != nil {
		return nil, errors.Wrap(err, "fail to send handshake")
	}

	clientChan := c.InterceptEvents(serverChan)

	return clientChan, nil
}

func (c *ChanClient) Do(actions *models.Actions) error {
	err := c.Send(models.MutationType, map[string]interface{}{"mutations": actions.Actions})
	if err != nil {
		return errors.Wrap(err, "fail to send actions")
	}
	return nil
}

func (c *ChanClient) Send(method string, payload interface{}) error {
	err := c.Transport.Send(models.Method{
		AgentID: c.ID,
		Method:  method,
		Payload: payload,
	})
	if err != nil {
		return errors.Wrap(err, "fail to send method")
	}
	return nil
}
