package main

import (
	"log"
	"time"
)

const (
	TimeUnit             = 2 * time.Second
	TableNumber          = 2
	WaiterNumber         = 1
	MaxFoods             = 6
	KitchenServerAddress = "http://localhost:8087/order"
	//PortNumber         = ":8086"
)

var (
	OrderNumber = 0
	//ranking     = 0.0
)

func main() {
	//initialising list of tables
	var TableList = make([]Table, 2*TableNumber)
	initTables(TableList)
	//initialising list of waiters
	var WaiterList = make([]Waiter, 2*WaiterNumber)
	initWaiters(WaiterList, TableList)

	go CheckTableState(TableList)
	//TODO:implement the server side of the DiningHALL
	for {
	}
	//var server = http.NewServeMux()
	//initServer(server, PortNumber)
}

func initTables(tList []Table) {
	for i := 0; i < TableNumber; i++ {
		tList[i].Init(i)
		log.Printf("initialising table #%v with state %v\n", tList[i].id, tList[i].state)
		go TableController(&tList[i])
	}
}

func initWaiters(wList []Waiter, tList []Table) {
	for i := 0; i < WaiterNumber; i++ {
		wList[i].Init(i)
		log.Printf("initialising waiter #%v with state %v\n", i, 0)
		go wList[i].Start(tList)
	}
}
func CheckTableState(tList []Table) {
	for {
		for i := 0; i < TableNumber; i++ {
			log.Printf("table %v; state %v", tList[i].id, tList[i].state)
		}
		time.Sleep(1 * TimeUnit)
	}
}
