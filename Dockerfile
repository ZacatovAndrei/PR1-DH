FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod .
COPY *.go ./
COPY menu.json .

RUN go build -o DiningHall

EXPOSE 8086/tcp

ENTRYPOINT ["/app/DiningHall"]