# Usage

```Bash
# ha-proxy
$ docker exec -it rabbitmq_clustered_golang_1 go run /app/mq_producer.go -u=admin -p=Passw0rd -h=localhost -t=0 -limit=50000 -msize=1024

# direct RabbitMQ
$ docker exec -it rabbitmq_clustered_golang_1 go run /app/mq_producer.go -u=admin -p=Passw0rd -h=localhost -t=0 -limit=50000 -msize=1024
```
