#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/Tools/kafka_2.13-3.0.0

# Start Kafka
$KAFKA_HOME/bin/kafka-server-start.sh $KAFKA_HOME/config/server.properties
