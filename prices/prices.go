package prices

import (
	"fmt"

	"example.com/price-calculator-v1/conversion"
	"example.com/price-calculator-v1/io_manager"
)

type TaxIncludePriceJob struct {
	IOManager io_manager.IOManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func NewTaxIncludePriceJob(fm io_manager.IOManager,taxRate float64) *TaxIncludePriceJob{
	return &TaxIncludePriceJob{
		IOManager: fm,
		TaxRate: taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}

func (job *TaxIncludePriceJob) Process() (error) {
	job.LoadDate()

	result := make(map[string]string)

	for _, price := range(job.InputPrices){
		priceWithTax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", priceWithTax)
	}

	job.TaxIncludedPrices = result
	err := job.IOManager.WriteLines(job)
	return err
}

func(job *TaxIncludePriceJob) LoadDate()(error){
	lines, err := job.IOManager.ReadLines()

	if err != nil{
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	
	if err != nil{
		return err
	}

	job.InputPrices = prices

	return nil
}