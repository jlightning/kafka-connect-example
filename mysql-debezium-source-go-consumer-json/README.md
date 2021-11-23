## Simple MySQL Debezium source with Go consumer JSON format

```
docker-compose up -d

# Wait a few minutes

# Run this command to register mysql source for kafka connect
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @register-mysql-source.json

go mod tidy
go run main.go
```

You will see go application log all changes made to `customers` table 
