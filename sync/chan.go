package sync

import "sync"

// Chan is structure type that will drop event if the client does not consume it fast enough.
// Once the client will consume an event, chan will ensure that this event is the last one
// Chan has received
type Chan struct {
	lock      sync.Mutex
	value     interface{}
	readyChan chan bool
}

// Push will add the event to the channel
func (c *Chan) Push(v interface{}) {
	c.lock.Lock()
	c.value = v
	c.lock.Unlock()

	select {
	case c.readyChan <- true:
	default:
	}
}

// Pop will return the last event that we received. If we didn't received any events since the last
// call, this function will block until a new event is received.
func (c *Chan) Pop() interface{} {
	<-c.readyChan
	c.lock.Lock()
	v := c.value
	c.lock.Unlock()
	return v
}

// NewChan initialize a new channel
func NewChan() *Chan {
	return &Chan{
		readyChan: make(chan bool),
	}
}
