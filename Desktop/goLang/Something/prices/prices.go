package prices

import "strconv"

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string][]float64
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 39},
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
