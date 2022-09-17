package main

import (
	"log"
	"time"
)

const (
	TimeUnit             = 2 * time.Second
	TableNumber          = 2
	WaiterNumber         = 1
	PortNumber           = ":8086"
	KitchenServerAddress = ":8087/order"
)

var (
	OrderNumber = 0
	ranking     = 0.0
)

type (
	tableArray  [TableNumber]Table
	WaiterArray [WaiterNumber]Waiter
)

func main() {
	var TableList tableArray
	initTables(&TableList)
	var WaiterList WaiterArray
	initWaiters(&WaiterList, &TableList)

	go CheckTableState(&TableList)
	//var server = http.NewServeMux()
	//initServer(server, PortNumber)
	println(time.Now().Unix())
	for {
	}
}

func initTables(tList *tableArray) {
	for i := 0; i < TableNumber; i++ {
		tList[i] = NewTable(i)
		log.Printf("initialising table #%v with status %v\n", tList[i].id, tList[i].status)
		go TableController(&tList[i])
	}
}

func initWaiters(wList *WaiterArray, tlist *tableArray) {
	for i := 0; i < WaiterNumber; i++ {
		wList[i] = NewWaiter(i)
		log.Printf("initialising waiter #%v with status %v\n", i, 0)
		go WaiterController(&wList[i], tlist)
	}
}
func CheckTableState(tList *tableArray) {
	for {
		for i := 0; i < TableNumber; i++ {
			log.Printf("table %v; state %v", tList[i].id, tList[i].status)
		}
		time.Sleep(1 * TimeUnit)
	}
}
