package agent

import (
	"os"

	"github.com/johnsudaar/go-bytearena/models"
	"github.com/johnsudaar/go-bytearena/transport"
	"github.com/pkg/errors"
)

type Agent struct {
	Transport transport.Transport
	ID        string
	Version   string
}

func FromEnv(version string) (*Agent, error) {
	if os.Getenv("AGENTID") == "" {
		return nil, errors.New("no agent id specified")
	}

	t, err := transport.FromEnv()
	if err != nil {
		return nil, err
	}

	a := &Agent{
		Transport: t,
		ID:        os.Getenv("AGENTID"),
		Version:   version,
	}

	return a, nil
}

func (a *Agent) Start() (chan interface{}, error) {
	c, err := a.Transport.Connect()
	if err != nil {
		return nil, err
	}

	c2 := a.InterceptEvents(c)

	return c2, nil
}

func (a *Agent) Do(actions models.Actions) error {
	err := a.Send(models.MutationType, map[string]interface{}{"mutations": actions})
	if err != nil {
		return errors.Wrap(err, "fail to send actions")
	}
	return nil
}

func (a *Agent) Send(method string, payload interface{}) error {
	err := a.Transport.Send(models.Method{
		AgentID: a.ID,
		Type:    method,
		Payload: payload,
	})
	if err != nil {
		return errors.Wrap(err, "fail to send method")
	}
	return nil
}
