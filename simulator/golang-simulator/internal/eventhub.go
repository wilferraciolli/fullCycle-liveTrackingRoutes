package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type EventHub struct {
	routeService    *RouteService
	mongoClient     *mongo.Client
	chDriverMoved   chan *DriverMovedEvent
	freightWriter   *kafka.Writer
	simulatorWriter *kafka.Writer
}

func NewEventHub(routeService *RouteService, mongoClient *mongo.Client, chDriverMoved chan *DriverMovedEvent, freightWriter *kafka.Writer, simulatorWriter *kafka.Writer) *EventHub {
	return &EventHub{
		routeService:    routeService,
		mongoClient:     mongoClient,
		chDriverMoved:   chDriverMoved,
		freightWriter:   freightWriter,
		simulatorWriter: simulatorWriter,
	}
}

// handler for events within Go
func (eh *EventHub) HandleEvent(msg []byte) error {
	// create base event for handling multiple types
	var baseEvent struct {
		EventName string `json:"event"`
	}
	// unmarshall message unto base event - used to determine event type by EventName field
	err := json.Unmarshal(msg, &baseEvent)
	if err != nil {
		return fmt.Errorf("error unmarshelling event : %w", err)
	}

	switch baseEvent.EventName {
	case "RouteCreated":
		var event RouteCreatedEvent
		err := json.Unmarshal(msg, &event)
		if err != nil {
			return fmt.Errorf("error unmarshelling event : %w", err)
		}

		return eh.handleRouteCreated(event)
	case "DeliveryStarted":
		var event DeliveryStartedEvent
		err := json.Unmarshal(msg, &event)
		if err != nil {
			return fmt.Errorf("error unmarshelling event : %w", err)
		}

		return eh.handleDeliveryStarted(event)
	default:
		return errors.New("unknown event")
	}
}

func (eh *EventHub) handleRouteCreated(event RouteCreatedEvent) error {
	// handle event
	freightCalculatedEvent, err := RouteCreatedHandler(&event, eh.routeService)
	if err != nil {
		return err
	}
	value, err := json.Marshal(freightCalculatedEvent)
	if err != nil {
		return fmt.Errorf("error marshalling event : %w", err)
	}

	// send event to Kafka
	err = eh.freightWriter.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(freightCalculatedEvent.RouteID),
		Value: value,
	})
	if err != nil {
		return fmt.Errorf("error writing message : %w", err)
	}

	return nil
}

func (eh *EventHub) handleDeliveryStarted(event DeliveryStartedEvent) error {
	// handle event
	err := DeliveryStartedHandler(&event, eh.routeService, eh.chDriverMoved)
	if err != nil {
		return err
	}

	// get driver moved from the channel driver moved on the handler.go then send event to Kafka, thi is done via GO creating a new routine
	go eh.sendDirections()

	return nil
}

func (eh *EventHub) sendDirections() {
	// make infinite look to keep listening to the channel. This will handle messages and have a timer to make itself idle
	for {
		select {
		case movedEvent := <-eh.chDriverMoved:
			value, err := json.Marshal(movedEvent)
			if err != nil {
				return
			}

			err = eh.simulatorWriter.WriteMessages(context.Background(), kafka.Message{
				Key:   []byte(movedEvent.RouteID),
				Value: value,
			})

			if err != nil {
				return
			}
		case <-time.After(500 * time.Millisecond):
		}
	}
}

// Class used to manage flow of events
// Nest.js -> routeTopic -> Go
// Go -> freight -> Nest.js
// Nest.js -> routeTopic -> Go (dlivery started)
// Go -> simulator -> Nest.js
