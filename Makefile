run:
	go run main.go waiter.go table.go order.go
build: 
	go build -o DiningHall main.go waiter.go table.go order.go
clean: 
	rm DiningHall
docker:
	docker build -t zacatov/pr1 .
