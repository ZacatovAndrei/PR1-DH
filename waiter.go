package main

import (
	"bytes"
	"container/list"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
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

func (w *Waiter) Start(tableList []Table, oList *list.List) {
	rand.Seed(time.Now().UnixNano())
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
		//however the reference to the initial OrderList seems to be lost
		//TODO: make taking responses from the list work
		log.Println("No orders found in the hall,checking kitchen")
		time.Sleep(TimeUnit)
		err := w.deliverOrder(oList, tableList)
		if err != nil {
			fmt.Println("error in delivery")
			continue
		}
		fmt.Println("Delivery successful")

	}
}
func (w *Waiter) deliverOrder(ol *list.List, tl []Table) error {
	//getting an order from the list of prepared orders
	if ol.Len() < 1 {
		return errors.New("no orders in the kitchen")
	}
	tOrder := ol.Front()
	w.CurrentOrder = tOrder.Value.(*Order)
	fmt.Printf("%v", w.CurrentOrder)
	ol.Remove(tOrder)
	log.Printf("took order:%v\n", w.CurrentOrder)

	//deliver
	if id := w.CurrentOrder.TableId; id >= TableNumber {
		log.Printf("No table found with id %v", id)
	}
	tl[w.CurrentOrder.TableId].state = done
	Rank = Rank + tl[w.CurrentOrder.TableId].rank(w.CurrentOrder)
	CompletedOrders++
	return nil
}

func (w *Waiter) takeOrder(table *Table) {
	//incrementing the global number of orders
	OrderNumber++
	//needed so that the waiters know where and which order has to be delivered
	table.orderID = OrderNumber
	//generating a new order
	numFoods := rand.Intn(MaxFoods) + 1
	items := make([]int, 0)
	log.Printf("%v/%v", numFoods, MaxFoods)
	for i := 0; i < numFoods; i++ {
		items = append(items, rand.Intn(13))
	}
	//getMaxPrepTime(items)
	w.CurrentOrder = NewOrder(OrderNumber, table.id, w.id, items, rand.Intn(5)+1, 45, time.Now().Unix())
	log.Printf("Order from table #%v taken by waiter #%v\n", table.id, w.id)
}

func (w *Waiter) sendOrder(order *Order, address string) {
	var b []byte
	b, ok := json.Marshal(order)
	if ok != nil {
		log.Fatalln("Could not Marshal JSON")
	}
	if resp, err := http.Post(address, "text/json", bytes.NewBuffer(b)); err != nil {
		fmt.Printf("%v", resp)
		panic(err)
	}
	w.CurrentOrder = nil
}
