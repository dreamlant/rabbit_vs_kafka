#!/bin/bash -

docker exec rabbitmq_clustered_rabbitmq-01_1 rabbitmqctl set_policy policy_ha_all "" '{"ha-mode":"all","ha-sync-mode":"automatic"}'

docker exec -it rabbitmq_clustered_rabbitmq-02_1 rabbitmqctl stop_app
docker exec -it rabbitmq_clustered_rabbitmq-02_1 rabbitmqctl join_cluster rabbit@rabbitmq-01
docker exec -it rabbitmq_clustered_rabbitmq-02_1 rabbitmqctl start_app

docker exec -it rabbitmq_clustered_rabbitmq-03_1 rabbitmqctl stop_app
docker exec -it rabbitmq_clustered_rabbitmq-03_1 rabbitmqctl join_cluster rabbit@rabbitmq-01
docker exec -it rabbitmq_clustered_rabbitmq-03_1 rabbitmqctl start_app
