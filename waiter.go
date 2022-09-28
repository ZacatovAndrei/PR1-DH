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

var (
	listAccess sync.Mutex
	searchLock sync.Mutex
)

type Waiter struct {
	id           int
	CurrentOrder *Order
}

func (w *Waiter) Init(i int) {
	w.id = i
	w.CurrentOrder = nil
	log.Printf(cGreen+"initialising waiter #%v"+cResetNl, i)
}

func (w *Waiter) Start(id int, tableList []Table, oList *list.List) {
	rand.Seed(time.Now().UnixNano())
	w.Init(id)
	for {
		tid := w.searchForOrder(tableList)
		if tid >= 0 {
			w.takeOrder(&tableList[tid])
			time.Sleep(3 * TimeUnit)
			w.sendOrderToKitchen(w.CurrentOrder)
			time.Sleep(2 * TimeUnit)
		}
		err := w.deliverOrder(oList, tableList)
		if err == nil {
			fmt.Println(cGreen + "Delivery successful" + cReset)
		}
		time.Sleep(TimeUnit)
	}
}

func (w *Waiter) searchForOrder(tl []Table) int {
	rand.Seed(time.Now().UnixNano())
	ci := -1
	searchLock.Lock()
	defer searchLock.Unlock()
	for i := 0; i < TableNumber; i++ {
		//Randomising starting table location
		//looping over the slice starting at a random location
		ci = (rand.Intn(TableNumber) + i) % TableNumber
		if tl[ci].state == occupied {
			tl[ci].state = waiting
			return ci
		}
	}

	return -1
}

func (w *Waiter) takeOrder(table *Table) {
	atomic.AddInt64(&OrderNumber, 1)
	//generating a new order
	numFoods := rand.Intn(MaxFoods) + 1
	items := make([]int, numFoods)
	for i := 0; i < numFoods; i++ {
		items[i] = rand.Intn(13)
	}

	w.CurrentOrder = NewOrder(int(OrderNumber),
		table.id,
		w.id,
		items,
		rand.Intn(5)+1,
		0,
		time.Now().Unix())
	//TODO: incorporate max wait into orderCreation
	w.CurrentOrder.getMaxPrepTime()
	log.Printf(cCyan+"Order from table #%v taken by waiter #%v"+cResetNl, table.id, w.id)

}

func (w *Waiter) deliverOrder(ol *list.List, tl []Table) error {

	listAccess.Lock()
	//should unlock the mutex at the end of the function even if returns by error path
	defer listAccess.Unlock()
	if ol.Len() < 1 {
		//listAccess is unlocked here in case condition fails
		return errors.New("there are no orders in the kitchen")
	}
	tOrder := ol.Front()
	w.CurrentOrder = tOrder.Value.(*Order)
	ol.Remove(tOrder)
	log.Printf(cCyan+"took order:%v"+cResetNl, *w.CurrentOrder)
	//deliver
	tl[w.CurrentOrder.TableId].state = done
	Rank += int64(tl[w.CurrentOrder.TableId].rank())
	CompletedOrders++
	//listAccess is unlocked here with normal execution flow
	return nil
}

func (w *Waiter) sendOrderToKitchen(order *Order) {
	//serializing JSON
	var b []byte
	b, ok := json.Marshal(order)
	if ok != nil {
		log.Panicln(cRed + "Could not Marshal JSON" + cReset)
	}
	//POSTing it to the Kitchen server
	if resp, err := http.Post(KitchenServerAddress, "application/json", bytes.NewBuffer(b)); err != nil {
		fmt.Printf("response:\t%v", resp)
		panic(err)
	}
	//logs + cleaning the current order from waiter's memory
	log.Println(cGreen + "POST request succeeded" + cReset)
	w.CurrentOrder = nil
}
