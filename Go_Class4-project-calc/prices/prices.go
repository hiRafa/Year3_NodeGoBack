package prices

import (
	"fmt"

	"example.com/m/v2/conversion"
	"example.com/m/v2/readcontent"
)

type PricesWithTaxesJob struct {
	IOManager       readcontent.FileManager `json:"-"`
	TaxRate         float64                 `json:"tax_rate"`
	InputPrices     []float64
	PricesWithTaxes map[string]string
}

func NewPricesWithTaxesJob(fileM readcontent.FileManager, taxRate float64) *PricesWithTaxesJob {
	return &PricesWithTaxesJob{
		IOManager:   fileM,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}
}

func (job *PricesWithTaxesJob) LoadData() {
	// lines, error := readcontent.ReadContent("prices.txt")
	lines, error := job.IOManager.ReadContent()

	if error != nil {
		fmt.Println(error)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *PricesWithTaxesJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		pricesWithTax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.1f", pricesWithTax)
	}

	job.PricesWithTaxes = result
	job.IOManager.WriteJSON(job)
	// readcontent.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)
}
