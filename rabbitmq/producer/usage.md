$ go get github.com/streadway/amqp
$ go run mq_producer.go -u=admin -p=Passw0rd -h=localhost -t=0 -limit=10 -msize=5