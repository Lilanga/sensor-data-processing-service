package timescaledb

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Lilanga/sensor-data-processing-service/internal/models"
	"github.com/jackc/pgx/v4/pgxpool"
)

type TimescaleDBClient struct {
	pool *pgxpool.Pool
}

func GetTimescaleDBClient() *TimescaleDBClient {
	var client = &TimescaleDBClient{}
	client.init()

	return client
}

func (c *TimescaleDBClient) init() {
	ctx := context.Background()
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=require", os.Getenv("TSDB_USER"), os.Getenv("TSDB_PASS"), os.Getenv("TSDB_HOST"), os.Getenv("TSDB_PORT"), os.Getenv("TSDB_DB"))
	dbpool, err := pgxpool.Connect(ctx, connStr)

	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	c.pool = dbpool
}

func (c *TimescaleDBClient) InsertReadings(readings []models.Reading) {
	// Insert contents of readings slice into TimescaleDB
	for i := range readings {
		r := readings[i]
		err := c.InsertReading(r)
		if err != nil {
			log.Printf("Unable to insert sample into Timescale %v\n", err)
			os.Exit(1)
		}
	}
	log.Println("Successfully inserted samples into sensor_data hypertable")
}

func (c *TimescaleDBClient) InsertReading(reading models.Reading) error {
	queryInsertTimeseriesData := `
	INSERT INTO weather_readings (time, deviceid, temperature, humidity) VALUES ($1, $2, $3, $4);
	`
	ctx := context.Background()
	_, err := c.pool.Exec(ctx, queryInsertTimeseriesData, reading.Time, reading.SensorId, reading.Temperature, reading.Humidity)
	if err != nil {
		log.Printf("Unable to insert sample into Timescale %v\n", err)
		return err
	}

	return nil
}
