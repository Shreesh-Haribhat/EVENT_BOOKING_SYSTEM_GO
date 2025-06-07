package main

import (
	"fmt"

	p "shreesh.com.in/prices"
)

func main() {
	fmt.Println("Hello wellcome to my application")

	// prices := []float64{10, 20, 30}
	taxRate := []float64{0, 0.2, 0.3, 0.5}

	// result := make(map[float64][]float64)

	for _, taxRate := range taxRate {
		job1 := p.NewTaxIncludedPriceJob(taxRate)
		job1.Calculate()
		fmt.Println(job1)
	}

}
