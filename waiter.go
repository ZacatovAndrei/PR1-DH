package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Waiter struct {
	id           int
	CurrentOrder *Order
	OrderList    []Order
}

func (w *Waiter) Init(i int) {
	//TODO: add a new parameter for kitchen's prepared order list
	w.id = i
	w.CurrentOrder = nil
	w.OrderList = nil
}

func (w *Waiter) Start(tableList []Table) {
	for {
		//Searching for an order
		//Randomising starting table location
		var startTable, currentIndex int

		startTable = rand.Intn(TableNumber)
		//looping over the slice starting at a random location
		for i := 0; i < TableNumber; i++ {
			currentIndex = (startTable + i) % TableNumber
			log.Printf("approaching table #%v, state - %v", tableList[currentIndex].id, tableList[currentIndex].state)
			if tableList[currentIndex].state == occupied {
				//taking order
				w.takeOrder(&tableList[currentIndex])
				tableList[currentIndex].state = waiting
				//sleep for a bit because taking an order takes time
				time.Sleep(5 * TimeUnit)
				//send order to kitchen
				w.sendOrder(w.CurrentOrder, KitchenServerAddress)
			}
		}
		//if there are no tables to serve -> deliver orders from the kitchen
		//TODO: implement a list of orders received from the kitchen
		log.Println("No orders found,checking kitchen")
		time.Sleep(5 * TimeUnit)
		//if err:=w.deliverOrder(ol []Order,t []Table); err!=nil{
		//log.Printf("Waiter #%v has found no orders in the kitchen\n",w.id)
		//}

	}
}

func (w *Waiter) takeOrder(table *Table) {
	//incrementing the global number of orders
	OrderNumber++
	//needed so that the waiters know where and which order has to be delivered
	table.orderID = OrderNumber
	//generating a new order
	numFoods := rand.Intn(MaxFoods) + 1
	items := make([]int, 10)
	for i := 0; i < numFoods; i++ {
		items = append(items, rand.Intn(13))
	}
	//getMaxPrepTime(items)
	w.CurrentOrder = NewOrder(OrderNumber, table.id, w.id, items, rand.Intn(5)+1, 45, time.Now().Unix())
	log.Printf("Order from table #%v taken by waiter #%v\n", table.id, w.id)
}

func (w *Waiter) sendOrder(order *Order, address string) {
	b, ok := json.Marshal(order)
	fmt.Printf("Order:\n%v", string(b))
	if ok != nil {
		log.Fatalln("Could not Marshal JSON")
	}
	fmt.Println("fake POST request succeeded")
	//if resp, err := http.Post(address, "text/json", bytes.NewBuffer(b)); err != nil {
	//	fmt.Printf("%v", resp)
	//	panic(err)
	//}
}
