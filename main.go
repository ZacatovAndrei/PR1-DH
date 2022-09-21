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

//Constants for coloured output
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
	KitchenServerAddress = "http://0.0.0.0:8087/order"
	LocalAddress         = ":8086"
)

var ()

var (
	CompletedOrders int32      = 0
	OrderNumber     int32      = 0
	Rank            int64      = 0
	OrderList       *list.List = list.New()
)

func main() {
	//initialising list of tables
	var TableList = make([]Table, TableNumber)
	initTables(TableList)
	//initialising list of waiters
	var WaiterList = make([]Waiter, WaiterNumber)
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
		log.Printf(cGreen+"initialising waiter #%v\n"+cReset, i)
		go wList[i].Start(i, tList, oList)
	}
}
