# Description

Repository produces PoC for Go and PHP based producers and consumers to perform a benchmark and throughput testing by using underlying Go/PHP langs.   

## Usage

1. Build steps:
`$ build.sh`
Wait while containers will be up & running
`$ clusterize.sh`

2. Attack using Go producer
```Bash
docker exec -it rabbitmq_clustered_golang_1 go run /app/mq_producer.go -u=admin -p=Passw0rd -h=rabbitmq-01 -t=0 -limit=100000 -msize=1024
```

Open from browser HA-Proxy stat:
`http://localhost:1936/`

RabbitMQ Management:
`http://localhost:15672/`


# Results
## Go Producer
```
$ docker exec -it rabbitmq_clustered_golang_1 go run /app/mq_producer.go -u=admin -p=Passw0rd -h=rabbitmq-01 -t=0 -limit=100000 -msize=1024

2020/04/16 09:18:30 Connecting to amqp://admin:Passw0rd@rabbitmq-01:5672/
2020/04/16 09:18:30 Connected.
2020/04/16 09:18:37 Took 7.263858 seconds
2020/04/16 09:18:37 Throughut 13766.789159 m/s
```
## Go Consumer
```
$ docker exec -it rabbitmq_clustered_golang_1 go run /app/mq_consumer.go -u=admin -p=Passw0rd -h=rabbitmq-01

2020/04/16 09:20:41 Connecting to amqp://admin:Passw0rd@rabbitmq-01:5672/
2020/04/16 09:20:41 Connected.
2020/04/16 09:20:55 Took 14.151737 seconds
2020/04/16 09:20:55 Throughut 0.706627 m/sec
```

## PHP Producer using HA-Proxy
```
$ docker exec -it rabbitmq_clustered_php71-cli_1 php /app/mq_producer.php rabbitmq-01 admin Passw0rd 100000 1024

RabbitMQ connected
Message size: 1024 bytes.
RabbitMQ connected
Message size: 1024 bytes.
```

## PHP Producer using directly RabbitMQ connection
```
$ docker exec -it rabbitmq_clustered_php71-cli_1 php /app/php_producer.php rabbitmq-01 admin Passw0rd 100000 1024

Elapsed: 8.988051 seconds
Throughput: 11125.882652 / second
```