package main

import (
	"context"
	"fmt"

	"github.com/fullCycle-liveTrackingRoutes/simulator/internal"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	println("Hello world")
	// connect to DB
	mongoStr := "mongodb://admin:admin@localhost:27017/routes?authSource=admin"
	mongoConnection, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoStr))
	if err != nil {
		panic(err)
	}

	// instantiate services
	freightService := internal.NewFreightService()
	routeService := internal.NewRouteService(mongoConnection, freightService)

	routeCreatedEvent := internal.NewRouteCreatedEvent(
		"1",
		100,
		[]internal.Directions{{Lat: 1, Lng: 1}})

	fmt.Println(internal.RouteCreatedHandler(routeCreatedEvent, routeService))
	// TODO stoped at 1:19 minutes
}
