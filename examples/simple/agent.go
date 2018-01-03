package main

import (
	"fmt"

	"github.com/ByteArena/go-bytearena/agent"
	"github.com/ByteArena/go-bytearena/models"
)

type Agent struct {
	agent.EmptyAgent
}

func (*Agent) Tick(perception models.Perception, _ int) *models.Actions {
	actions := &models.Actions{}
	actions.Steer(models.NewVector2(0, 1))
	actions.Shoot(models.Vector2{0, 0})
	return actions
}

func (*Agent) Error(err error, _ int) {
	fmt.Printf("Error received: %s\n", err.Error())
}
