package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/fullCycle-liveTrackingRoutes/simulator/internal"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	println("Go simulator started")

	mongoURI := getEnv("MONGO_URI", "mongodb://admin:admin@mongo:27017/routes?authSource=admin")
	kafkaBroker := getEnv("KAFKA_BROKER", "kafka:9092")
	kafkaRouteTopic := getEnv("KAFKA_ROUTE_TOPIC", "route")
	kafkaFreightTopic := getEnv("KAFKA_FREIGHT_TOPIC", "freight")
	kafkaSimulationTopic := getEnv("KAFKA_SIMULATION_TOPIC", "simulation")
	kafkaGroupID := getEnv("KAFKA_GROUP_ID", "route-group")

	// connect to DB
	mongoConnection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}

	// instantiate services
	freightService := internal.NewFreightService()
	routeService := internal.NewRouteService(mongoConnection, freightService)

	// instantiate GO channels
	chDriverMoved := make(chan *internal.DriverMovedEvent)

	// instantiate kafka
	freightWriter := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    kafkaFreightTopic,
		Balancer: &kafka.LeastBytes{},
	}
	simulatorWriter := &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    kafkaSimulationTopic,
		Balancer: &kafka.LeastBytes{},
	}

	// creat event reader
	routeReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaBroker},
		Topic:   kafkaRouteTopic,
		GroupID: kafkaGroupID,
	})

	// create event handler
	hub := internal.NewEventHub(routeService, mongoConnection, chDriverMoved, freightWriter, simulatorWriter)

	// subscribe to kafka
	fmt.Println("Starting simulator")
	for {
		m, err := routeReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error: %w, err")
		}

		// create a go routine for new thread so it wont block it
		go func(msg []byte) {
			err = hub.HandleEvent(m.Value)
			if err != nil {
				log.Printf("error handling event: %w, err")
			}
		}(m.Value)
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
