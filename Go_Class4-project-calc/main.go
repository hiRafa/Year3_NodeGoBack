package main

import (
	"fmt"

	"example.com/m/v2/prices"
	"example.com/m/v2/readcontent"
)

func main() {
	// var prices []float64 = []float64{10,20,30}
	// prices := []float64{10, 20, 30}
	taxRates := []float64{0, 0.7, 0.1, 0.15, 0.02}
	// results := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		fileM := readcontent.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewPricesWithTaxesJob(fileM, taxRate)
		priceJob.Process()
	}
	// for _, taxRate := range taxRates {
	// 	pricesWithTaxes := make([]float64, len(prices))
	// 	for priceIndex, price := range prices {
	// 		pricesWithTaxes[priceIndex] = price * (1 + taxRate)
	// 	}
	// 	results[taxRate] = pricesWithTaxes
	// }
	// fmt.Println(results)
}
