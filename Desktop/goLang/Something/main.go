package main

import (
	"fmt"

	p "shreesh.com.in/prices"
)

func main() {
	fmt.Println("Hello wellcome to my application")

	taxRate := []float64{0, 0.2, 0.3, 0.5}

	for _, taxRate := range taxRate {
		job1 := p.NewTaxIncludedPriceJob(taxRate)
		job1.GetPricesFromFile()
		job1.Calculate()
		fmt.Println(job1)
	}

}
