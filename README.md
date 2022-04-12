### Sensor data processing service

This sample service is written for reading writing from HiveMQ private mqtt cloud endpoint.

Use appropriate values for following environment values
```
// MQTT Host URL in 'tls://xxxxxxx....' format for private HiveMQ endpoints
MQTT_HOST=

// MQTT port number
MQTT_PORT=

// MQTT user name
MQTT_USER=

// MQTT password
MQTT_PASS=

// string to be used as the client for the connection
MQTT_CLIENT_ID=

// Timescale Database user
TSDB_USER=

// Timescale Database password
TSDB_PASS=

// Timescale database host url
TSDB_HOST=

// Timescale database name
TSDB_DB=

// Port of the timescale database
TSDB_PORT=

// SSL mode of the Timescaledb connection. one of `require`, `prefer`, `allow`, `disabled`
TSDB_SSL_MODE=
```

#### Docker build
You can use `dockerBuild.sh` file inside the build folder to create docker image for testing