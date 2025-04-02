#!/bin/bash

CONFIG_FILE="/etc/openmeter.yaml"

echo "🔍 Looking for config at $CONFIG_FILE"
if [[ -f "$CONFIG_FILE" ]]; then
  echo "✅ Found config. Launching OpenMeter with it."
  exec /usr/local/bin/openmeter --config "$CONFIG_FILE"
else
  echo "❌ Config file not found. Exiting."
  exit 1
fi

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
