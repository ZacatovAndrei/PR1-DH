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
	"sync"
	"sync/atomic"
	"time"
)

type Waiter struct {
	id           int
	CurrentOrder *Order
}

func (w *Waiter) Init(i int) {
	w.id = i
	w.CurrentOrder = nil
}

func (w *Waiter) Start(id int, tableList []Table, oList *list.List) {
	rand.Seed(time.Now().UnixNano())
	w.Init(id)
	for {
		tid := w.searchForOrder(tableList)
		if tid >= 0 {
			w.takeOrder(&tableList[tid])
			time.Sleep(3 * TimeUnit)
			w.sendOrderToKitchen(w.CurrentOrder, KitchenServerAddress)
			time.Sleep(2 * TimeUnit)
		}
		if tid == -1 {
			log.Println(cYellow + "No orders found in the hall,checking kitchen" + cReset)
		}
		err := w.deliverOrder(oList, tableList)
		if err != nil {
			log.Println(cYellow, err, cReset)
		} else {
			fmt.Println(cGreen + "Delivery successful" + cReset)
		}
		time.Sleep(TimeUnit)
	}
}

func (w *Waiter) searchForOrder(tl []Table) int {
	var searchLock sync.Mutex
	rand.Seed(time.Now().UnixNano())
	ci := -1
	searchLock.Lock()
	defer searchLock.Unlock()
	for i := 0; i < TableNumber; i++ {
		//Randomising starting table location
		//looping over the slice starting at a random location
		ci = (rand.Intn(TableNumber) + i) % TableNumber
		//log.Printf("approaching table #%v, state - %v", tl[ci].id, tl[ci].state)
		if tl[ci].state == occupied {
			return ci
		}
	}

	return -1
}

func (w *Waiter) takeOrder(table *Table) {
	atomic.AddInt32(&OrderNumber, 1)
	//generating a new order
	numFoods := rand.Intn(MaxFoods) + 1
	items := make([]int, numFoods)
	for i := 0; i < numFoods; i++ {
		items[i] = rand.Intn(13)
	}
	//getMaxPrepTime(items)
	w.CurrentOrder = NewOrder(int(OrderNumber), table.id, w.id, items, rand.Intn(5)+1, 45, time.Now().Unix())
	log.Printf(cCyan+"Order from table #%v taken by waiter #%v\n"+cReset, table.id, w.id)
	table.state = waiting

}

func (w *Waiter) deliverOrder(ol *list.List, tl []Table) error {
	var listAccess sync.Mutex

	listAccess.Lock()
	//should unlock the mutex at the end of the function even if returns by error path
	defer listAccess.Unlock()
	if ol.Len() < 1 {
		return errors.New("there are no orders in the kitchen")
	}
	tOrder := ol.Front()
	w.CurrentOrder = tOrder.Value.(*Order)
	ol.Remove(tOrder)
	log.Printf(cCyan+"took order:%v\n"+cReset, *w.CurrentOrder)
	//deliver
	tl[w.CurrentOrder.TableId].state = done
	Rank += tl[w.CurrentOrder.TableId].rank(w.CurrentOrder)
	CompletedOrders++
	return nil
}

func (w *Waiter) sendOrderToKitchen(order *Order, address string) {
	var b []byte
	b, ok := json.Marshal(order)
	if ok != nil {
		log.Panicln(cRed + "Could not Marshal JSON" + cReset)
	}
	if resp, err := http.Post(address, "application/json", bytes.NewBuffer(b)); err != nil {
		fmt.Printf("response:\t%v", resp)
		panic(err)
	}
	log.Println(cGreen + "POST request succeeded" + cReset)
	w.CurrentOrder = nil
}
