# Simulator - GoLang

used to simulate the cars moving around the map. The main usage for this service is to handle and publish events Eg carMoved, routeCreated...
Microservice built using the Go language.

# Prerequisites
Install Go Lang from here [https://go.dev/dl/](https://go.dev/dl/)
Then install the VsCode extensions to manage Go
* Go extension (from official GO)
* plugins as per picture below. to install the plugins, press ctrl+shift+p and type `GO: Install/Update tools`
* Ps it will output on the console all the tools being installed, if fails then run again

![Install GO plugin and tools.png](../images/image1-installGOAndTools.png)

# Optional plugins
### MongoDB explorer
This is used to connect to the mongoDB running on a container for SQL queries. This is a VSCode plugin made by Mongo, it is the official SQL editor.
![Install MongoDB extension.png](../images/image3-mongoDBExtension.png)
To connect with it, simply pass in the connection string defined within the environment variables `mongodb://admin:admin@localhost:27017/routes?authSource=admin` 
this will allow to query and see the data inserted onto the database.

# Initialize Go project (Create initial module)
within the terminal type in `go mod init github.com/fullCycle-liveTrackingRoutes/simulator` 
this will create the initial module pointing to the github and the version of GO used. this address can be used for downloading the microservice


# Running the main function
on the command line, type in `go run cmd/simulator/main.go`


# Running the main function within docker compose
on the command line, do a `docker compose up -d` then withinthe console type in `docker compose exec -it simulator sh` this will get the 
GoLang `simulator` container and open the bash within it. Finally then we can run the normal command to start the app `go run cmd/simulator/main.go`

# Kafka Control Center (GUI)
This is used to monitor the messages coming in and out of kafka, ideal for debuging.
To access it, there is a container on `localhost:9021` which is connected to kafka to allow it

# Downloading dependencies
On the command line, type in `go mod tidy` as seen image below

![Install dependencies.png](../images/image2-installDependencies.png)


# Dependencies
### Apache Kafka GO
Used to publish and receive message within Kafka
`github.com/segmentio/kafka-go`

### Mongo DB
Database `go.mongodb.org/mongo-driver`
