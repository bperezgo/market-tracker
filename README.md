# MARKET-TRACKER

## Description of the project

It is used a pool connections of websocket, or another information source, to get market data. Can be used any kind of interfaces to get this data. For the moment only websockets, all this data is processed and sent to a data management system, but with a specific structure, to mantain data consistency.

In other words, this market tracker application is a data minner.

But it does not matter if only take data of different sources, the important thing is doing something with this data. So, all the gathered data will be sent to a system with ML or other system to analyze the data (In a next iteration)

It will be used the next [websocket tool](https://github.com/nhooyr/websocket) because is light, and accomplish all the needs of this project, at least to take the data, provisioned in the websoket apis.

## Setup the project

If you want to run the tests, you need this command

```bash
docker run -d --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic events.dummy.tested --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181
```

### Init the project

The file configuration_example.json is an example of how it is necesary to setup the environment variables for this project. To run this project, it must be created a file, that is called configuration.json, in the root of the project, in the same part of the configuration_example.json. Then, change the dummy values for its respective real values.

#### Development

Ensure to install [conventional-pre-commit](https://github.com/compilerla/conventional-pre-commit)
