package agent

import (
	"encoding/json"
	"fmt"

	"github.com/johnsudaar/go-bytearena/models"
	"github.com/pkg/errors"
)

func (a *Agent) InterceptEvents(c chan interface{}) chan interface{} {
	res := make(chan interface{})
	go func() {
		for {
			raw := <-c
			switch v := raw.(type) {
			case error:
				res <- models.ErrorEvent{Error: v}
			case []byte:
				a.RouteEvent(v, res)
			default:
				panic(raw)
			}
		}
	}()
	return res
}

func (a *Agent) RouteEvent(v []byte, c chan interface{}) {
	c <- models.RawEvent{
		Value: v,
	}

	var e models.Event
	err := json.Unmarshal(v, &e)
	if err != nil {
		c <- models.ErrorEvent{Error: errors.Wrap(err, "fail to unmarshal server event")}
		return
	}

	switch e.Method {
	case models.MethodPerception:
		var perception models.Perception
		err := json.Unmarshal(e.Payload, &perception)
		if err != nil {
			c <- models.ErrorEvent{Error: errors.Wrap(err, "fail to unmarshal perception")}
			return
		}
		c <- perception
	case models.MethodWelcome:
		var welcome models.Welcome
		err := json.Unmarshal(e.Payload, &welcome)
		if err != nil {
			c <- models.ErrorEvent{Error: errors.Wrap(err, "fail to unmarshal welcome")}
			return
		}
		c <- welcome
	default:
		c <- models.ErrorEvent{Error: errors.New(fmt.Sprintf("Unknown event: %s", e.Method))}
	}
}
