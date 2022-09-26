package main

import "time"

//Constants for coloured output
const (
	cReset   = "\033[0m"
	cResetNl = cReset + "\n" // simplifies one line log.printf logs
	cRed     = "\033[31m"
	cGreen   = "\033[32m"
	cYellow  = "\033[33m"
	cBlue    = "\033[34m"
	cPurple  = "\033[35m"
	cCyan    = "\033[36m"
	cGray    = "\033[37m"
	cWhite   = "\033[97m"
)

//global configuration variables
const (
	TimeUnit                        = 2 * time.Second
	TableNumber                     = 10
	WaiterNumber                    = 4
	MaxFoods                        = 6
	KitchenServerAddress            = "http://Kitchen:8087/order"
	KitchenServerAddressNoContainer = "http://localhost:8087/order"
	LocalAddress                    = ":8086"
	MenuPath                        = "./"
)
