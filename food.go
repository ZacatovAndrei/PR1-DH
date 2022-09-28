package main

import (
	"encoding/json"
	"io/ioutil"
)

type RestaurantMenu []FoodItem

type FoodItem struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PreparationTime  int    `json:"preparation-time"`
	Complexity       int    `json:"complexity"`
	CookingApparatus string `json:"cooking-apparatus"`
}

func (m RestaurantMenu) ParseMenu(s string) RestaurantMenu {
	fin, err := ioutil.ReadFile(s)
	if err != nil {
		panic(err)
	}
	var menu RestaurantMenu
	if err := json.Unmarshal(fin, &menu); err != nil {
		panic(err)
	}
	return menu

}
