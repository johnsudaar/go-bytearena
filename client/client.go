package client

import (
	"strings"

	"github.com/johnsudaar/go-bytearena/agent"
	"github.com/johnsudaar/go-bytearena/models"
	"github.com/johnsudaar/go-bytearena/sync"
	"github.com/pkg/errors"
)

type Client struct {
	Agent          agent.Agent
	internalChan   chan interface{}
	chanClient     *ChanClient
	Specs          models.Specs
	workerLaunched bool
	endChan        chan bool
}

func New() (*Client, error) {
	chanClient, err := NewChanClient(VersionClearBeta)
	if err != nil {
		return nil, errors.Wrap(err, "fail to create client")
	}

	client := &Client{
		chanClient: chanClient,
		endChan:    make(chan bool),
	}

	return client, nil
}

func (c *Client) Start(a agent.Agent) error {
	internalChan, err := c.chanClient.Start()
	if err != nil {
		return errors.Wrap(err, "fail to start client")
	}
	c.internalChan = internalChan
	c.Agent = a

	go c.Worker()
	<-c.endChan
	return nil
}

func (c *Client) Do(acts *models.Actions) error {
	return c.chanClient.Do(acts)
}

func (c *Client) Worker() {
	errorChan := sync.NewChan()
	perceptionChan := sync.NewChan()
	rawChan := sync.NewChan()

	go c.PerceptionWorker(perceptionChan)
	go c.ErrorWorker(errorChan)
	go c.RawWorker(rawChan)

	for {
		raw := <-c.internalChan
		switch v := raw.(type) {
		case models.ErrorEvent:
			if strings.Contains(v.Error.Error(), "EOF") {
				c.endChan <- true
			} else {
				errorChan.Push(v.Error)
			}
		case models.Specs:
			if !c.workerLaunched {
				c.Specs = v
				go c.Agent.Worker()
			}
		case models.Perception:
			perceptionChan.Push(v)
		case models.RawEvent:
			rawChan.Push(v)
		default:
		}
	}
}

func (c *Client) PerceptionWorker(perceptionChan *sync.Chan) {
	for {
		actions := c.Agent.Tick(perceptionChan.Pop().(models.Perception), 0)
		if actions != nil && len(actions.Actions) > 0 {
			c.chanClient.Do(actions)
		}
	}
}

func (c *Client) ErrorWorker(errorChan *sync.Chan) {
	for {
		c.Agent.Error(errorChan.Pop().(error), 0)
	}
}

func (c *Client) RawWorker(rawChan *sync.Chan) {
	for {
		c.Agent.Raw(rawChan.Pop().(models.RawEvent).Value, 0)
	}
}
