run:
	go run main.go waiter.go table.go order.go
build: 
	go build -o DiningHall 
clean: 
	rm DiningHall
docker:
	if [ -n $(docker image ls | grep zacatov/pr1dining)]; then docker rmi zacatov/pr1dining;fi
	docker build -t "zacatov/pr1dining" .

