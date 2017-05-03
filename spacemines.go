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
	c := Colony{}
	c.numMines = random(3,6)
	c.numPeople = random(40, 60)
	c.money = random(10, 50) * c.numPeople
	c.foodPrice = random(40,80)
	c.oreProduction = random(40,80)

	c.oreStorage = 0
	c.year = 1;
	c.satisfaction = 1

	c.rollPriceDice()

	c.failed = false
	return &c;
}

func (c *Colony) rollPriceDice() {
	c.minePrice = random(2000, 4000)
	c.orePrice = random(7, 12)
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func randomFloat() float64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Float64()
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

func askForIntInput(s string) int {
	var output int
	for {
		fmt.Printf("%s", s)
		_, err := fmt.Scanf("%d\n",&output)

		if err != nil {
			fmt.Println("That input was invalid")
			continue
		} else {
			break
		}
	}
	return output
}

func (c *Colony) oreSale() {
	for {
		oreToSell := askForIntInput("How much ore to sell? ")
		if oreToSell >= 0 && oreToSell <= c.oreStorage{
			c.oreStorage -= oreToSell
			c.money += oreToSell * c.orePrice
			break
		}
	}
}

func (c *Colony) mineSale() {
	for {
		minesToSell := askForIntInput("How many mines to sell? ")
		if minesToSell >= 0 && minesToSell <= c.numMines{
			c.numMines -= minesToSell
			c.money += minesToSell * c.minePrice
			break
		}
	}
}

func (c *Colony) foodBuy() {
	for {
		foodToBuy := askForIntInput("How much to spend on food? (Appr. $100 EA.) ")
		if foodToBuy >= 0 && foodToBuy <= c.money{
			c.food += foodToBuy
			c.money -= foodToBuy

			if foodToBuy / c.numPeople > 120 {
				c.satisfaction+=.1
			}

			if foodToBuy / c.numPeople < 80 {
				c.satisfaction-=.2
			}
			break
		}else{
			fmt.Println("You don't have enough money to afford that amount of food.")
		}
	}
}

func (c *Colony) mineBuy() {
	for {
		minesToBuy := askForIntInput("How many more mines to buy? ")
		if minesToBuy >= 0 && (minesToBuy * c.minePrice) <= c.money{
			c.numMines += minesToBuy
			c.money -= minesToBuy * c.minePrice
			break
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

		// If there are less than 10 people per mine then game over
		if c.numPeople / c.numMines < 10 {
			c.failed = true
			fmt.Println("You've overworked everyone, Game Over!")
			break
		}

		// If satisfaction is high, more people arrive
		if c.satisfaction > 1.1 {
			c.numPeople += random(1,10)
		}

		// People leave if satisfaction is low
		if c.satisfaction < 0.9 {
			c.numPeople -= random(1,10)
		}

		// If the satisfaction is too low then game over
		if c.satisfaction < 0.6 {
			c.failed = true
			fmt.Println("The people revolted, Game Over!")
			break
		}

		// If there are less than 30 people in total then game over
		if c.numPeople < 30 {
			c.failed = true
			fmt.Println("Not enough people left, Game Over!")
			break
		}

		// Introduce a small chance that half the population gets killed
		if randomFloat() < 0.1 {
			fmt.Println("RADIOACTIVE LEAK....MANY DIE!")
			c.numPeople /= 2
		}

		// If the amount produced per mine is very high, ore price is halved
		if c.oreProduction > 150 {
			fmt.Println("Market Glut - Price Drops!")
			c.foodPrice /= 2
		}

		// Player has survived another year
		c.rollPriceDice()
		c.year++
		fmt.Println("")
	}

	if c.failed == false {
		fmt.Println("You survived your term of office")
	}
}