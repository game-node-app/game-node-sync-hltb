# game-node-sync-hltb

The GameNode sync system for HLTB.

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
After installing the project dependencies, simply use these commands to start each service:  
```shell

go run cmd/listener/main.go

go run cmd/processor/main.go
```

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

