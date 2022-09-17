package main

type Order struct {
	OrderId        int              `json:"order_id,omitempty"`
	TableId        int              `json:"table_id,omitempty"`
	WaiterId       int              `json:"waiter_id,omitempty"`
	Items          []int            `json:"items,omitempty"`
	Priority       int              `json:"priority,omitempty"`
	MaxWait        int              `json:"max_wait,omitempty"`
	PickUpTime     int64            `json:"pick_up_time,omitempty"`
	CookingTime    int              `json:"cooking_time,omitempty"`
	CookingDetails []map[string]int `json:"cooking_details,omitempty"`
}

func NewOrder(orderId, tableId, waiterId int, items []int, priority, maxWait int, pickUpTime int64) *Order {
	return &Order{
		OrderId:    orderId,
		TableId:    tableId,
		WaiterId:   waiterId,
		Items:      items,
		Priority:   priority,
		MaxWait:    maxWait,
		PickUpTime: pickUpTime}
}
