run:
	go run main.go waiter.go table.go order.go
build: 
	go build -o DiningHall main.go waiter.go table.go order.go food.go
clean: 
	rm DiningHall
docker:
	docker rmi zacatov/pr1dining ;
	docker build -t "zacatov/pr1dining" .

