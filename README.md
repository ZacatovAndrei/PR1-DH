# PR1-DH

This repository is the home to the first part of the laboratory work \#1 for the Network programming course at UTM, the second one being the [Kitchen](https://github.com/zahatikoff/PR1-KT)

## The purpose
The purpose of the Dining Hall is to create orders that will be delivered to the kitchen with the help of the waiters, and to receive the already prepared orders from the kitchen to the clients, and based off of the timings, calculating the real time ranking of the restaurant simulation.

Kitchen opens a server on `http://localhost:8086` and accepting POST requests on the `/distribution` path 
### Available commands
Some commands are available from a makefile as a simplification:
- `make build` - builds an executable `Kitchen`
- `make docker` - creates a docker container with the name zacatov/pr1dining

### Running a container: 
Containers should be run on the same docker network that one can create via a `docker network create`,
since it allows using the names of the containers as their names as their IP addresses,
so the sample command would look like 
`docker run --name DiningHall --network restaurant zacatov/pr1dining`