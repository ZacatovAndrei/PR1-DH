package main

import (
	"log"
	"math/rand"
	"time"
)

const (
	free int = iota
	occupied
	waiting
	done
)

type Table struct {
	id, status int
}

func NewTable(index int) Table {
	log.Printf("Initialising table %v \n", index)
	retTab := Table{status: free, id: index}
	return retTab
}
func TableController(table *Table) {
	var random, oldState int
	for {
		switch table.status {
		case free:
			random = rand.Intn(6) + 10
			time.Sleep(time.Duration(random) * TimeUnit)
			oldState = table.status
			table.status = occupied
			log.Printf("status of table %v changed from %v to %v;Table occupied", table.id, oldState, table.status)
		case done:
			random = rand.Intn(16) + 15
			time.Sleep(time.Duration(random) * TimeUnit)
			oldState = table.status
			table.status = free
			log.Printf("Order has been served;status of table %v changed from %v to %v;Table freed", table.id, oldState, table.status)
		}
	}
}
