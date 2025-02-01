package internal

import "time"

type RouteCreatedEvent struct {
	EventName  string       `json:"event"`
	RouteID    string       `json:"id"`
	Distance   int          `json:"distance"`
	Directions []Directions `json:"directions"`
}

func NewRouteCreatedEvent(routeID string, distance int, directions []Directions) *RouteCreatedEvent {
	return &RouteCreatedEvent{
		EventName:  "RouteCreated",
		RouteID:    routeID,
		Distance:   distance,
		Directions: directions,
	}
}

type FreightCaculatedEvent struct {
	EventName string  `json:"event"`
	RouteID   string  `json:"route_id"`
	Amount    float64 `json:"amount"`
}

func NewFreightCalculatedEvent(routeId string, amount float64) *FreightCaculatedEvent {
	return &FreightCaculatedEvent{
		EventName: "FreightCalculated",
		RouteID:   routeId,
		Amount:    amount,
	}
}

type DeiveryStartedEvent struct {
	EventName string `json:"event"`
	RouteID   string `json:"route_id"`
}

func NewDeliveryStartedEvent(routeID string) *DeiveryStartedEvent {
	return &DeiveryStartedEvent{
		EventName: "DeliveryStarted",
		RouteID:   routeID,
	}
}

type DriverMovedEvent struct {
	EventName string  `json:"event"`
	RouteID   string  `json:"route_id"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}

func NewDriverMovedEvent(routeID string, lat, lng float64) *DriverMovedEvent {
	return &DriverMovedEvent{
		EventName: "DriverMoved",
		RouteID:   routeID,
		Lat:       lat,
		Lng:       lng,
	}
}

func RouteCreatedHandler(event *RouteCreatedEvent, routeService *RouteService) (*FreightCaculatedEvent, error) {
	route := NewRoute(event.RouteID, event.Distance, event.Directions)
	routeCreated, err := routeService.CreateRoute(route)
	if err != nil {
		return nil, err
	}

	FreightCaculatedEvent := NewFreightCalculatedEvent(routeCreated.ID, routeCreated.FreightPrice)

	return FreightCaculatedEvent, nil
}

func DeliveryStartedHandler(event *DeiveryStartedEvent, routeService *RouteService, ch chan *DriverMovedEvent) error {
	route, err := routeService.GetRoute(event.RouteID)
	if err != nil {
		return err
	}

	// start go routine for async values - used to stream lat anf lng
	driverMovedEvent := NewDriverMovedEvent(route.ID, 0, 0)
	for _, direction := range route.Directions {
		driverMovedEvent.RouteID = route.ID
		driverMovedEvent.Lat = direction.Lat
		driverMovedEvent.Lng = direction.Lng
		// add delay to simulate driver movement
		time.Sleep(time.Second)
		// create a channel to share the driverMovedEvent to avoid concurrency issues
		ch <- driverMovedEvent
	}

	return nil
}
