package main

import (
	"fmt"

	"github.com/ByteArena/go-bytearena/client"
	"github.com/ByteArena/go-bytearena/models"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.TextFormatter{ForceColors: true, FullTimestamp: true})
	log.Info("Starting...")

	client, err := client.NewChanClient(client.VersionClearBeta)
	if err != nil {
		log.WithError(err).Panic("fail to init agent")
	}

	c, err := client.Start()
	if err != nil {
		log.WithError(err).Panic("fail to start agent")
	}

	for {
		raw := <-c
		switch v := raw.(type) {
		case models.ErrorEvent:
			log.Panic("Error received: ", v.Error.Error())
		case models.Specs:
			log.Info("Handshake response received")
		case models.Perception:
			log.Info("Perception received")
			acts := &models.Actions{}
			acts.Steer(models.Vector2{0, 1})
			log.Info("Sending actions")
			err := client.Do(acts)
			if err != nil {
				log.WithError(err).Panic("fail to send actions")
			}
		case models.RawEvent:
			log.Debug("Raw event received")
		default:
			log.Warn(fmt.Sprintf("Unexpected event type: %T", raw))
		}
	}
}
