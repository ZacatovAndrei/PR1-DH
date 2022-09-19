package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	TimeUnit             = 2 * time.Second
	TableNumber          = 2
	WaiterNumber         = 1
	MaxFoods             = 6
	KitchenServerAddress = "http://localhost:8087/order"
	LocalAddress         = "localhost:8086"
)

var (
	OrderNumber           = 0
	Rank            int64 = 0
	CompletedOrders       = 0
	OrderList             = list.New()
)

func main() {
	//initialising list of tables
	var TableList = make([]Table, 2*TableNumber)
	initTables(TableList)
	//initialising list of waiters
	var WaiterList = make([]Waiter, 2*WaiterNumber)
	initWaiters(WaiterList, TableList, OrderList)
	go CheckTableState(TableList)
	//TODO:implement the server side of the DiningHALL

	http.HandleFunc("/distribution", getOrder)
	if err := http.ListenAndServe(LocalAddress, nil); err != nil {
		panic(err)
	}

}

func getOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "The server only supports POST requests\n")
		return
	}
	b, ok := ioutil.ReadAll(r.Body)
	if ok != nil {
		panic(ok)
	}
	o := new(Order)
	if err := json.Unmarshal(b, o); err != nil {
		panic(err)
	}
	OrderList.PushFront(o)
	log.Printf("there are %v orders in the List now", OrderList.Len())
}

func initTables(tList []Table) {
	for i := 0; i < TableNumber; i++ {
		tList[i].Init(i)
		log.Printf("initialising table #%v with state %v\n", tList[i].id, tList[i].state)
		go TableController(&tList[i])
	}
}

func initWaiters(wList []Waiter, tList []Table, oList *list.List) {
	for i := 0; i < WaiterNumber; i++ {
		wList[i].Init(i)
		log.Printf("initialising waiter #%v with state %v\n", i, 0)
		go wList[i].Start(tList, oList)
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
