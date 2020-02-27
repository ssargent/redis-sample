package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ssargent/redis-sample/data"
)

// LoginSession is...
type LoginSession struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func main() {
	fmt.Println("redissample - test app")

	myDataService, _ := data.NewDataService("localhost", 6379)
	elapsed := []int64{}

	// store 10k keys
	for i := 0; i < 10000; i++ {

		session := LoginSession{Id: uuid.New(), Name: fmt.Sprintf("foo%d", i)}

		start := time.Now()
		myDataService.AddRecord("foo", session)
		duration := time.Since(start)

		elapsedTime := int64(duration / time.Millisecond)

		if elapsedTime > 250 {
			fmt.Printf("Record %d took %dms \n", i, elapsedTime)
		}

		elapsed = append(elapsed, elapsedTime)
	}

	min, max := MinMax(elapsed)

	var accum int64

	for i := 0; i < len(elapsed); i++ {
		accum += elapsed[i]
	}

	average := accum / int64(len(elapsed))

	fmt.Printf("Min %d  Max %d  Average %d\n", min, max, average)
}

func MinMax(array []int64) (int64, int64) {
	var max int64 = array[0]
	var min int64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
