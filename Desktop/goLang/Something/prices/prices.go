package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string][]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{},
	}
}

func (t *TaxIncludedPriceJob) Calculate() {
	mp := make(map[string][]float64)
	temp := []float64{}

	for _, val := range t.InputPrices {
		temp = append(temp, val*t.TaxRate)
	}

	mp[strconv.FormatFloat(t.TaxRate, 'f', 2, 64)] = temp
	t.TaxIncludedPrices = mp

}

func (t *TaxIncludedPriceJob) GetPricesFromFile() {
	filePtr, err := os.Open("prices.txt")
	if err != nil {
		fmt.Println("something went wrong")
		return
	}

	scanner := bufio.NewScanner(filePtr)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Failed to fetch the file content")
		return
	}

	for _, val := range lines {
		num, _ := strconv.ParseFloat(val, 64)
		t.InputPrices = append(t.InputPrices, num)
	}
	// fmt.Println(t.InputPrices)

}
