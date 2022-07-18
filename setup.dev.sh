#!/bin/sh

# Remember change the persmissions of this file
chmod +x ./pipelines/**/*.sh
pre-commit install
# TODO: Used to initialize the dev project, and work with the dev config in local

# This command works ok initializing the first topic, over this, will be iterated
# docker exec -it f2fc3c95a6e5 /opt/bitnami/kafka/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --topic events.asset.recorded --partitions 3 --replication-factor 2