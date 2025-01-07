# Nest js backend server
The back end will be used to get places from google maps, create directions from Google maps and finally to create a route within the database.
It integrates with a database and will be our REST Webb server
![Nest js back end 1.png](../images/1-overview.png)

The server is modular and will have different modules created to build the backend
![2-nest-js-architecture.png](../images/2-nest-js-architecture.png)

# Installing Nest Js CLI
run `npm install -g @nestjs/cli`

### Creating the project
run the command `nest new nestjs-api` and chose the `npm` as the package manager
This will create a node project with a package .json

### Running the app
run `npm run start:dev` to spin the server, this should create a dummy hello world endpoint on the address 
[http://localhost:300](localhost:300)
![3-server-run.png](../images/3-server-run.png)

### Running the app with DOCKER
run `docker-compose up` this command will deploy both the app and the DB server, each of them should have a corresponding Dockerfile
if an image already exists, then it uses the current image, otherwise make sure to build the latest docker image before running the app,
alternativetely, you can pass the --build flag force the build to happen rather than using the current image Eg `docker-compose up --build`

once docker compose is up, then on the terminal type in `docker compose exec app bash` this will allow you to see what the content of the service called `app` is
very useful to check the content of what the container has, simply type `ls` on the container bash session and it will display its content or mannually starting the app eg `npm `
![5-conatiner-bash.png](../images/5-conatiner-bash.png)

## Nest JS CLI commands
### Creating Modules
run `nest g module <moduleName>` to create a new module, this will create an empty module and add it as imports within the main module

### Creating controller
run `nest g controller <moduleName/controllerName>`

### Creating service
run `nest g service <moduleName/serviceName>`

### Creating CRUD resources
run `nest g resource`

## Testing within the IDE
Creating a file called `fileName.http` will allow the ide to create a local http client and perform requests
![4-test-http-client.png](../images/4-test-http-client.png)


# Dependencies
### Google maps
Google maps are used to query places, calculate routes and display
`npm i @googlemaps/google-maps-services-js` 
Google maps will need the following API and Services enabled
* Places API
* Directions API
* Others needed for the UI to display map, check ui project

### Prisma ORM
Used to map ts classes onto database tables and vice vers
`npm i @prisma/client`
The dependency above requires the connection to be set up on start up of the module, 
the service can implement the OnModuleInit interface and connect.
```typescript
@Injectable()
export class PrismaService extends PrismaClient implements OnModuleInit {
  async onModuleInit() {
    await this.$connect();
  }
}
```
**IMPORTANT** The commands above, especially the prisma generate, must be run within the docker exec bash terminal, as it OS agnostic,
hence doing it is Windows whilst developing will not work within Docker, it will send the message of 
```angular2html
PrismaClientInitializationError: Prisma Client could not locate the Query Engine for runtime "debian-openssl-3.0.x".

This happened because Prisma Client was generated for "windows", but the actual deployment required "debian-openssl-3.0.x".

```

**then**, within the project root run the following to instantiate Prisma within the project `npx prisma init`, 
this will install and instantiate the project as prisma enabled. A few folders called prisma will be added, 
some variables and config, including adding the connect string to the .env file.
Note that within the prisma folder created (next to node_modules), there is the schema.prisma file which is the main file for config

**then**, within the project root run `npx prisma generate` so it will validate and generate the schemas that needs to be synched to the database as well as install any other dependency needed


### Config - used to inject environment variables 
Used to inject environment variables Eg Google maps api
`npm install @nestjs/config`
Usage, within the main module, import and set it to be used globally, then to be used within the app, inject the ConfigService and do a get
```typescript
    @Module({
    imports: [
    // add config module from nestjs to be able to inject environment variables
        ConfigModule.forRoot({
        isGlobal: true
    }),
      MapsModule
    ],
    controllers: [AppController],
    providers: [AppService],
    })
export class AppModule {}

...
this.configService.get<string>('GOOGLE_MAPS_API_KEY')
  

```

# Database
MongoDB is used within this app
the database is run on a docker container, this specific instance of MongoDB 
is a bit different as it is configured as a replica within the Dockerfile on the .docker/mogodb folder

to connect with the database, the following string is needed 
`mongodb://root:root@localhost:27017/nest?authSource=admin` Ps note that `root`is both the username and password, 
then `nest` is the database that was created within the MongoDB dockerfile


# Function for terminal run NestJS
Called `repl`, it is really just a TS file that we can create Eg `repl.ts` which them we can specify the module we want to boot
and it will boot the module and open a command line where we can then call methods and pass parameters, this is really useful
for when you want to call a method that does not have an endpoint Eg add Points to the route-driver
```typescript
import { repl } from '@nestjs/core';
import { AppModule } from './app.module';

async function bootstrap() {
  await repl(AppModule);
}
bootstrap();
```
Then to run it, simply add the `entryFiile` attribute to the command line Eg `npm run start:dev -- --entryFile=repl`
And make sure to add the `await` keyword before calling methods otherwise it will fail silently.
Once the terminal runs, you can then do a `get()` on services and call their methods
![6-repl-cli.png](../images/6-repl-cli.png)












