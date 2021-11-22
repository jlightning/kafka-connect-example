## Simple MySQL Debezium source to JDBC Sink data synchronization with AVRO

```
docker-compose up -d

# Wait a few minutes

# Run this command to register mysql source for kafka connect
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @register-mysql-source.json

# Run this command to register mysql sink for kafka connect
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @register-mysql-sink.json
```

You will see data being synchronized between `customers` and `addresses` to `replicated_customers` and `replicated_addresses` when you connect to mysql on `127.0.0.1:3306`
