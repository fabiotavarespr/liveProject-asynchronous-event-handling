#!/bin/bash

# Set this to the location where Kafka has been installed
KAFKA_HOME=~/Tools/kafka_2.13-3.0.0

# Start Zookeeper
$KAFKA_HOME/bin/zookeeper-server-start.sh $KAFKA_HOME/config/zookeeper.properties