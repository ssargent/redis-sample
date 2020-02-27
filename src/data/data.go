package data

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

type DataService interface {
	AddRecord(key string, value interface{}) error
}

type dataService struct {
	server string
	port   int
}

func NewDataService(server string, port int) (DataService, error) {
	ds := &dataService{server: server, port: port}

	return ds, nil
}

// AddRecord adds a record.
func (d *dataService) AddRecord(key string, value interface{}) error {
	now := time.Now()
	score := now.UnixNano()
	conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", d.server, d.port))

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer conn.Close()

	record, _ := json.Marshal(value)
	_, err = conn.Do("ZADD", key, score, string(record))

	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
