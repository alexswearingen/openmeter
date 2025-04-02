#!/bin/bash

echo "=================================="
echo "🚀 Starting OpenMeter"
echo "Kafka Brokers: $KAFKA_BROKERS"
echo "Kafka Username: $KAFKA_USERNAME"
echo "Kafka Protocol: $KAFKA_SECURITY_PROTOCOL"
echo "Kafka Mechanism: $KAFKA_SASL_MECHANISM"
echo "Database URL: $DATABASE_URL"
echo "Redis URL: $REDIS_URL"
echo "=================================="

# Export all explicitly for safety
export KAFKA_BROKERS
export KAFKA_USERNAME
export KAFKA_PASSWORD
export KAFKA_SECURITY_PROTOCOL
export KAFKA_SASL_MECHANISM
export DATABASE_URL
export REDIS_URL
export OPENMETER_TOKEN

exec /usr/local/bin/openmeter
