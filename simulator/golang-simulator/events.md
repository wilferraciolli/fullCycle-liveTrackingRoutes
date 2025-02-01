# List of events that will be handled by the microservice
Here is the definition of events that the application will handle


# Recive event
RouteCreated
- id
- distance
- directions
-- lat
--lng

### Execute and return events Eg calculate cost
FreightCalculated
- route_id
- ammount

---
# Receive event
DeliveryStarted
- route_id

### Execute and return event
DriverMoved
-route_id
- lat
- lng













