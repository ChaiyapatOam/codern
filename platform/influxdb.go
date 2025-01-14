package platform

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxDb struct {
	client   influxdb2.Client
	writeApi api.WriteAPIBlocking
	queryApi api.QueryAPI
}

func NewInfluxDb(url string, token string, org string, bucket string) (*InfluxDb, error) {
	client := influxdb2.NewClient(url, token)
	writeApi := client.WriteAPIBlocking(org, bucket)

	if ok, err := client.Ping(context.Background()); !ok {
		return nil, err
	}

	return &InfluxDb{
		client:   client,
		writeApi: writeApi,
		queryApi: client.QueryAPI(org),
	}, nil
}

func (db *InfluxDb) WritePoint(
	measurement string,
	tags map[string]string,
	fields map[string]interface{},
) error {
	point := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	return db.writeApi.WritePoint(context.Background(), point)
}

func (db *InfluxDb) Close() {
	db.client.Close()
}
