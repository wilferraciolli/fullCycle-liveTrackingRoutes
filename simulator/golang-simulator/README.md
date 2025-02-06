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
To connect with it, simply pass in the connection string defined within the main file `mongodb://admin:admin@localhost:27017/routes?authSource=admin` 
this will allow to query and see the data inserted onto the database.

# Initialize Go project (Create initial module)
within the terminal type in `go mod init github.com/fullCycle-liveTrackingRoutes/simulator` 
this will create the initial module pointing to the github and the version of GO used. this address can be used for downloading the microservice


# Running the main function
on the command line, type in `go run cmd/simulator/main.go`

# Downloading dependencies
On the command line, type in `go mod tidy` as seen image below

![Install dependencies.png](../images/image2-installDependencies.png)


# Dependencies
### Apache Kafka GO
Used to publish and receive message within Kafka
`github.com/segmentio/kafka-go`
