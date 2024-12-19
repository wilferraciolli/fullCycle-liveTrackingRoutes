# fullCycle-liveTrackingRoutes
NextJs, NestJs, Go... technology to track routing in real time


Tutorial from the Full Cycle youtube channel on using NextJS, NestJS, Go, Kafka and Google Maps to create a full stack program for tracking cars.


# codebase
`https://github.com/devfullcycle/imersao20`

# Overview
* Backend - NestJs
* Frontend - NextJs, React, Tailwind
* Simulator - Golang
* Messaging - Kafka
* DB - MongoDB and Prisma ORM
* Messaging - Kafka

# 3rd party
Google Maps APIs
Websockets

# Logic
The Front end will be used to create routes
The backend will process the API request, and send an event to kafka
Kafka will then publish the event to the Simulador (Golang) to create the rota Eg from-to destinations

Next

Golang will create the route and send the message to kafka which will pusblish to the backend (NestJs)
Golang and Kafka will provide real time tracking websockets via websockets between 
frontend (NextJs) and backend (NextJs) to update the tracking in real time
