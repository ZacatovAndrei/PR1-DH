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
	cReset  = "\033[0m"
	cRed    = "\033[31m"
	cGreen  = "\033[32m"
	cYellow = "\033[33m"
	cBlue   = "\033[34m"
	cPurple = "\033[35m"
	cCyan   = "\033[36m"
	cGray   = "\033[37m"
	cWhite  = "\033[97m"
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
	log.Printf(cBlue+"there are %v orders in the List now"+cReset, OrderList.Len())
}

func initTables(tList []Table) {
	for i := 0; i < TableNumber; i++ {
		tList[i].Init(i)
		go TableController(&tList[i])
	}
}

func initWaiters(wList []Waiter, tList []Table, oList *list.List) {
	for i := 0; i < WaiterNumber; i++ {
		wList[i].Init(i)
		log.Printf(cGreen+"initialising waiter #%v with state %v\n"+cReset, i, 0)
		go wList[i].Start(tList, oList)
	}
}
