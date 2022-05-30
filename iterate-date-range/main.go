package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	startDate, err := time.Parse("2006-01-02", "2022-01-01")
	if err != nil {
		log.Println(err)
	}
	endDate, err := time.Parse("2006-01-02", "2022-01-28")
	if err != nil {
		log.Println(err)
	}

	for date := startDate; date.AddDate(0, 0, -1).Before(endDate); date = date.AddDate(0, 0, 1) {
		fmt.Println(date)
	}
}
