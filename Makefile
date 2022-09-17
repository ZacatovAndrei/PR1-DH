run:
	go run main.go waiter.go table.go order.go
build: 
	go build main.go waiter.go table.go order.go -o DiningHall
clean: 
	rm DiningHall

