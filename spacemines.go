package main

import (
	"fmt"
	"time"
	"math/rand"
)


type game struct {
	numMines int
	numPeople int
	money int
	foodPrice int
	oreProduction int
	oreStorage int
	year int
	satisfaction int
	minePrice int
	orePrice int
}

func initGame() *game {
	g := game{}
	g.numMines = random(3,6)
	g.numPeople = random(40, 60)
	g.money = random(10, 50) * g.numPeople
	g.foodPrice = random(40,80)
	g.oreProduction = random(40,80)
	return &g;
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func main(){
	g := initGame();
	fmt.Println(g)
}