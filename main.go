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

var (
	CompletedOrders int64      = 0
	OrderNumber     int64      = 0
	Rank            int64      = 0
	OrderList       *list.List = list.New()
	CurrentMenu     RestaurantMenu
)

func main() {
	//parsing menu
	CurrentMenu = CurrentMenu.ParseMenu(MenuPath + "menu.json")
	log.Printf("current menu :\n %+v\n", CurrentMenu)

	//initialising list of tables
	var TableList = make([]Table, TableNumber)
	initTables(TableList)

	//initialising list of waiters
	var WaiterList = make([]Waiter, WaiterNumber)
	initWaiters(WaiterList, TableList, OrderList)

	//initialising the server side
	http.HandleFunc("/distribution", getOrder)
	if err := http.ListenAndServe(LocalAddress, nil); err != nil {
		panic(err)
	}
	defer log.Printf("the rank is %v")
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprintf(w, "The server only supports POST requests\n")
		return
	}

	//reading the response body
	b, ok := ioutil.ReadAll(r.Body)
	if ok != nil {
		panic(ok)
	}

	//deserializing the Order object
	o := new(Order)
	if err := json.Unmarshal(b, o); err != nil {
		panic(err)
	}

	//locking listAccess so that only one thread could access the list
	//( the other threads being for example waiters' deliverOrder )
	listAccess.Lock()
	OrderList.PushFront(o)
	log.Printf(cBlue+"there are %v orders in the List now"+cResetNl, OrderList.Len())
	listAccess.Unlock()
}

func initTables(tList []Table) {
	for i := 0; i < TableNumber; i++ {
		go tList[i].Start(i)
	}
}

func initWaiters(wList []Waiter, tList []Table, oList *list.List) {
	for i := 0; i < WaiterNumber; i++ {
		go wList[i].Start(i, tList, oList)
	}
}

func CheckTableStates(tl []Table) {
	for {
		for i, table := range tl {
			fmt.Printf("table %v with state %v", i, table.state)
		}
		time.Sleep(TimeUnit)
	}
}
