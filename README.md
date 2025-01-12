# game-node-sync-hltb

The GameNode sync system for HLTB.

# Deprecated
This repository is here for historical purposes. We have abandoned the idea of HowLongToBeat syncing because of their constant attempts to difficult API querying.  
We understand their reasons and respect their decision, and as such, this repository is no longer maintained.

## Usage
This service is dependent on [game-node-server](https://github.com/game-node-app/game-node-server), since the RabbitMQ 
instance is configured there. Have it set-up before proceeding.

## Installation
#### Pre-requisites
- At least `go 1.23.1` in your machine.
- A `redis` instance running in port `9112` 
You can also change the redis addr (host:port) used by setting the `REDIS_ADDR` env variable.
- A `rabbitmq` instance running in port `localhost:5672`  
You can also change the RabbitMQ URL (default: `amqp://localhost:5672`) by setting the `REDIS_URL` env variable.  
If possible, have [game-node-server](https://github.com/game-node-app/game-node-server) set up before this.

### Starting
After installing the project dependencies (`go mod download`), simply use these commands to start each service:  
```shell
# Starts the 'listener' process
go run cmd/listener/main.go

# Starts the 'processor' process
go run cmd/processor/main.go
```

### Using the Docker image
You can also use this service by pulling and running the `lamarcke/game-node-sync-hltb` image.  
```shell
# Starts the 'listener' process with the 'hltb-listener' container name
docker run --name hltb-listener lamarcke/game-node-sync-hltb /app/listener

# Starts the 'processor' process with the 'hltb-processor' container name
docker run --name hltb-processor lamarcke/game-node-sync-hltb /app/processor
```
See the `docker-compose.prod.yml` for a compose file example.

### Implementation
This Go app has two main entrypoints:
- `listener`
- `processor`

Both are available at the `cmd` directory.

`listener` is responsible for receiving requests for playtime info updates, and creates new tasks based on said requests.
`processor` is a [`asynq`](https://github.com/hibiken/asynq) worker pool. It handles the actual processing of requests, and
publishes the results accordingly.

RabbitMQ is used as the message broker for this application, and it's instance is available as a dependency 
in our [main api](https://github.com/game-node-app/game-node-server).

The actual processing is quite simple:  
`update request received` -> `handled as task` -> `parse game name to improve results` 
-> `search for best possible match` -> `publishes result`

The process then continues on our main api, where data is made available to users.

## Disclaimer

GameNode is not affiliated with HowLongToBeat or it's partner websites/services in any way. 

We're doing our best to keep the requests WAY below a possible block threshold by having a single worker processing requests
and artificial timeouts in place. Requests to HLTB are only actually done every 4-8 seconds, and only when strictly necessary.

Data received from HLTB is also credited everytime it's shown in our client apps.

If you are a HLTB affiliate and wish to contact us, please feel free to email us here:
[support@gamenode.app](mailto:support@gamenode.app)

