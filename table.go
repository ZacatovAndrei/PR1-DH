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
	id, state, orderID int
}

func (t *Table) Init(i int) {
	t.id = i
	t.state = free
	log.Printf("Initialising table %v \n", i)
}

func (t *Table) Occupy() {
	t.state = occupied
	log.Printf("Table %v is now occupied", t.id)
}

func (t *Table) Free() {
	t.state = free
	log.Printf("The order #%v has been served,Table %v is now free", t.orderID, t.id)
}

func TableController(table *Table) {
	var random int
	for {
		switch table.state {
		case free:
			//it takes 10-15 units of time to occupy a new table
			random = rand.Intn(6) + 10
			time.Sleep(time.Duration(random) * TimeUnit)
			table.Occupy()
		case done:
			//it takes 15-20 units of time for people to leave after getting the order
			random = rand.Intn(6) + 15
			time.Sleep(time.Duration(random) * TimeUnit)
			table.Free()
		}
	}
}
