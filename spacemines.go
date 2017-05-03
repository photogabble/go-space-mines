package main

import (
	"fmt"
	"time"
	"math/rand"
)


type Colony struct {
	numMines int
	numPeople int
	money int
	food int
	foodPrice int
	oreProduction int
	oreStorage int
	year int
	satisfaction float32
	minePrice int
	orePrice int
	failed bool
}

func initColony() *Colony {
	g := Colony{}
	g.numMines = random(3,6)
	g.numPeople = random(40, 60)
	g.money = random(10, 50) * g.numPeople
	g.foodPrice = random(40,80)
	g.oreProduction = random(40,80)

	g.oreStorage = 0
	g.year = 1;
	g.satisfaction = 1

	g.minePrice = random(2000, 4000)
	g.orePrice = random(7, 12)

	g.failed = false
	return &g;
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func (c *Colony) displayColonyStats() {
	fmt.Println("Year", c.year)
	fmt.Println("There are", c.numPeople, "people in the colony")
	fmt.Println("You have", c.numMines, "mines and $", c.money)
	fmt.Println("Satisfaction Factor ", c.satisfaction)
	fmt.Println("")
	fmt.Println("Your mines produced ", c.oreProduction, "tons each")

	c.oreStorage += c.oreProduction * c.numMines

	fmt.Println("Ore in store:", c.oreStorage, "tons")
}

func (c *Colony) oreSale() {
	finished := false
	for finished == false {
		var oreToSell int
		fmt.Print("How much ore to sell? ")
		fmt.Scanf("%d",&oreToSell)
		if (oreToSell >= 0 && oreToSell <= c.oreStorage){
			c.oreStorage -= oreToSell
			c.money += (oreToSell * c.orePrice)
			finished = true
		}
	}
}

func (c *Colony) mineSale() {
	finished := false
	for finished == false {
		var minesToSell int
		fmt.Print("How many mines to sell? ")
		fmt.Scanf("%d",&minesToSell)
		if (minesToSell >= 0 && minesToSell <= c.numMines){
			c.numMines -= minesToSell
			c.money += (minesToSell * c.minePrice)
			finished = true
		}
	}
}

func (c *Colony) foodBuy() {
	finished := false
	for finished == false {
		var foodToBuy int
		fmt.Print("How much to spend on food? (Appr. $100 EA. ")
		fmt.Scanf("%d",&foodToBuy)
		if (foodToBuy >= 0 && foodToBuy <= c.money){
			c.food += foodToBuy
			c.money -= foodToBuy

			if (foodToBuy / c.numPeople > 120) {
				c.satisfaction+=.1
			}

			if (foodToBuy / c.numPeople < 80) {
				c.satisfaction-=.2
			}
			finished = true
		}
	}
}

func (c *Colony) mineBuy() {
	finished := false
	for finished == false {
		var minesToBuy int
		fmt.Print("How many more mines to buy? ")
		fmt.Scanf("%d",&minesToBuy)
		if (minesToBuy >= 0 && (minesToBuy * c.minePrice) <= c.money){
			c.numMines += minesToBuy
			c.money = (minesToBuy * c.minePrice)
			finished = true
		}
	}
}

func main(){
	c := initColony()

	for c.year <= 10 && c.failed == false {
		c.displayColonyStats()

		// Selling
		fmt.Println("Selling:")
		fmt.Println("Ore selling price: $", c.orePrice, "/ton")
		fmt.Println("Mine selling price: $", c.minePrice, "/mine")

		c.oreSale()
		c.mineSale()

		// Buying
		fmt.Println("")
		fmt.Println("You have $", c.money)
		fmt.Println("")
		fmt.Println("Buying")
		c.foodBuy()
		c.mineBuy()

		// If satisfaction is high, more people arrive
		if c.satisfaction > 1.1 {
			c.numPeople += random(1,10)
		}

		// People leave if satisfaction is low
		if c.satisfaction < 0.9 {
			c.numPeople -= random(1,10)
		}

		// If there are less than 30 people then game over
		if c.numPeople < 30 {
			c.failed = true
			fmt.Println("Not enough people left, game over!")
			break
		}

		if (c.failed == false) {
			c.year++
		}
	}
	//fmt.Println(g)
}