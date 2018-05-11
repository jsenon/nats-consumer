[![CircleCI](https://circleci.com/gh/jsenon/nats-consumer.svg?style=svg)](https://circleci.com/gh/jsenon/nats-consumer)
[![Go Report Card](https://goreportcard.com/badge/github.com/jsenon/nats-consumer)](https://goreportcard.com/report/github.com/jsenon/nats-consumer)

# Nats Consumer
 
This is a Simple nats consumer

### Func

Use container in order to debug nats 

## Build

Build locally
```
make 
```

## Run Container

Run Container
```
docker run -d -e MY_NATSBOOTSTRAP="123.123.123.123:9092" -e MY_TOPIC="kikoo"  jsenon/nats-consumer:latest
```

## Run on K8S Cluster

Deploy POD
```
kubectl apply -f k8s/deployment.yml
```

## Env

- MY_NATSBOOTSTRAP: Define nats bootstrap server ie.127.0.0.1
- MY_TOPIC: Define topic to check ie.mytest

## Todo

