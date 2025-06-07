package main

import (
	"fmt"

	Interface "shreesh.com.in/interface"
	"shreesh.com.in/players"
)

func main() {
	fmt.Println("Hello wellcome to my application")

	// taxRate := []float64{0, 0.2, 0.3, 0.5}

	// for _, taxRate := range taxRate {
	// 	job1 := p.NewTaxIncludedPriceJob(taxRate)
	// 	job1.GetPricesFromFile()
	// 	job1.Calculate()
	// 	fmt.Println(job1)
	// }

	// player := players.NewCricketPlayer("Shreesh", 34)
	player := players.NewFootBallPlayer("Shreesh", 34)

	doSomething(player)

}

func doSomething(allrounder Interface.Allrounder) {
	allrounder.Batting("shreesh", 25)

	allrounder.Bowling("Shrusti", 25)
}
