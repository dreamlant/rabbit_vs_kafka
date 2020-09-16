# Usage
### Using native AMQP extension 
```Bash
 # ha-proxy
$ docker exec -it rabbitmq_clustered_php71-cli_1 php /app/mq_producer.php localhost admin Passw0rd 50000 1024

# direct RabbitMQ
$ docker exec -it rabbitmq_clustered_php71-cli_1 php /app/mq_producer.php rabbitmq_clustered_rabbitmq-01_1 admin Passw0rd 50000 1024
```
Arguments list:
1. Rabbit entrypoint hostname/ip (HA-Proxy balancer or RabbitMQ direct IP could be used)
1. RabbitMQ username
1. RabbitMQ password
1. Batch size
1. Message size in bytes
