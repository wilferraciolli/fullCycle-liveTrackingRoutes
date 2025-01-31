package internal

import (
	"math"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Directions struct {
	Lat float64
	Lng float64
}

type Route struct {
	ID           string
	Distance     int
	Directions   []Directions
	FreightPrice float64
}

type FreightService struct {
}

func (fs *FreightService) calculate(distance int) float64 {
	return math.Floor((float64(distance)*0.15+0.3)*100) / 100
}

type RouteService struct {
	mongo          *mongo.Client
	freightService *FreightService
}

func (rs *RouteService) CreateRoute(route Route) (Route, error) {
	route.FreightPrice = rs.freightService.calculate(route.Distance)

	update := bson.M{
		"$set": bson.M{
			"distance":      route.Distance,
			"directions":    route.Directions,
			"freight_price": route.FreightPrice,
		},
	}

	filter := bson.M{"_id": route.ID}
	opts := options.Update().SetUpsert(true)

	_, err := rs.mongo.Database("routes").Collection("routes").UpdateOne(nil, filter, update, opts)
	if err != nil {
		return Route{}, err
	}

	return route, err
}

func (rs *RouteService) GetRoute(id string) (Route, error) {
	var route Route
	filter := bson.M{"_id": id}
	err := rs.mongo.Database("routes").Collection("routes").FindOne(nil, filter).Decode(&route)

	if err != nil {
		return Route{}, err
	}

	return route, nil
}

// TODO stoped lesson 4 at 45 minutes
