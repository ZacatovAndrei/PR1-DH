package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
)

type Waiter struct {
	Id           int
	CurrentOrder *Order
}

func NewWaiter(index int) Waiter {
	return Waiter{Id: index, CurrentOrder: nil}
}

func WaiterController(waiter *Waiter, tableList *tableArray) {
	for {
		//searching for an order

		// adding some randomised access to the tables
		startTable := rand.Intn(TableNumber)
		for i := 0; i < TableNumber; i++ {

			log.Printf("approaching table #%v with status %v\n", tableList[(startTable+i)%TableNumber].id, tableList[(startTable+i)%TableNumber].status)
			if tableList[(startTable+i)%TableNumber].status == occupied {
				//taking order
				takeOrder(waiter, &tableList[(startTable+i)%TableNumber])
				tableList[(startTable+i)%TableNumber].status = waiting
				log.Printf("Found order at table #%v by waiter #%v;Sending to the kitchen.\n", tableList[(startTable+i)%TableNumber].id, waiter.Id)
				//sleep for a bit
				time.Sleep(3 * TimeUnit)
				//send order to kitchen
				sendOrder(waiter.CurrentOrder, KitchenServerAddress)
			} else {
				log.Println("No order found")
				time.Sleep(2 * TimeUnit)
			}
		}
	}
}

func takeOrder(waiter *Waiter, table *Table) {
	println(OrderNumber)
	OrderNumber++
	waiter.CurrentOrder = NewOrder(
		OrderNumber,
		table.id,
		waiter.Id,
		[]int{1, 2, 3, 4, 5},
		10,
		10,
		time.Now().Unix(),
	)
	println(OrderNumber)
	log.Printf("Order from table #%v taken by waiter #%v\n", table.id, waiter.Id)
}

func sendOrder(order *Order, address string) {
	b, ok := json.Marshal(order)
	fmt.Printf("Order:\n%v", string(b))
	if ok != nil {
		log.Fatalln("Could not Marshal JSON")
	}
	fmt.Println("fake POST request succeeded")

	//http.Post(address+"/order", "text/json", bytes.NewBuffer(b))
}
