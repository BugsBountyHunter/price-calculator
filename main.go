package main

import (
	"fmt"

	"example.com/price-calculator-v1/file_manager"
	"example.com/price-calculator-v1/prices"
)

func main(){
		taxRats := []float64{0, 0.07, 0.1, 0.15}
		
		for _, taxRate := range(taxRats){
			fm := file_manager.New("prices.txt", fmt.Sprintf("result_%.2f.json", taxRate * 100))
			// cmd := cmd_manager.New()
			job := prices.NewTaxIncludePriceJob(fm, taxRate)
			job.Process()
		}

}