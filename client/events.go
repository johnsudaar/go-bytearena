package client

import (
	"encoding/json"
	"fmt"

	"github.com/ByteArena/go-bytearena/models"
	"github.com/pkg/errors"
)

func (c *ChanClient) InterceptEvents(serverChan chan interface{}) chan interface{} {
	clientChan := make(chan interface{})
	go func() {
		for {
			raw := <-serverChan
			switch v := raw.(type) {
			case error:
				clientChan <- models.ErrorEvent{Error: v}
			case []byte:
				c.RouteEvent(v, clientChan)
			default:
				panic(raw)
			}
		}
	}()
	return clientChan
}

func (c *ChanClient) RouteEvent(v []byte, clientChan chan interface{}) {
	clientChan <- models.RawEvent{
		Value: v,
	}

	var e models.Event
	err := json.Unmarshal(v, &e)
	if err != nil {
		clientChan <- models.ErrorEvent{Error: errors.Wrap(err, "fail to unmarshal server event")}
		return
	}

	switch e.Method {
	case models.MethodPerception:
		var perception models.Perception
		err := json.Unmarshal(e.Payload, &perception)
		if err != nil {
			clientChan <- models.ErrorEvent{Error: errors.Wrap(err, "fail to unmarshal perception")}
			return
		}
		clientChan <- perception
	case models.MethodWelcome:
		var specs models.Specs
		err := json.Unmarshal(e.Payload, &specs)
		if err != nil {
			clientChan <- models.ErrorEvent{Error: errors.Wrap(err, "fail to unmarshal welcome")}
			return
		}
		clientChan <- specs
	default:
		clientChan <- models.ErrorEvent{Error: errors.New(fmt.Sprintf("Unknown event: %s", e.Method))}
	}
}
