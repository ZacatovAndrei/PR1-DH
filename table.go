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
	id, state int
}

func (t *Table) Init(i int) {
	t.id = i
	t.state = free
	log.Printf(cGreen+"initialising table #%v with state %v\n"+cReset, t.id, t.state)
}

func (t *Table) Occupy() {
	t.state = occupied
	log.Printf(cBlue+"Table %v is now occupied"+cReset, t.id)
}

func (t *Table) Free() {
	t.state = free
	log.Printf(cBlue+"Table %v is now free"+cReset, t.id)
}

func (t *Table) rank(o *Order) int64 {
	return 5
}

func TableController(table *Table) {
	var random int
	for {
		switch table.state {
		case free:
			//it takes 5-10 units of time to occupy a new table
			random = rand.Intn(6) + 5
			time.Sleep(time.Duration(random) * TimeUnit)
			table.Occupy()
		case done:
			//it takes 5-10 units of time for people to leave after getting the order
			random = rand.Intn(6) + 5
			time.Sleep(time.Duration(random) * TimeUnit)
			table.Free()
		}
	}
}
