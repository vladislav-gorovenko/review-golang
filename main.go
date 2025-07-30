package main

import (
	"github.com/vladislav-gorovenko/review-golang/bmi_calculator"
)

func main() {
	//var convertAmount = converter.GetAmountInput()
	//fromCurrency, toCurrency := converter.GetCurrencyInput()
	//converter.ConvertCurrency(fromCurrency, toCurrency, &convertAmount)
	//fmt.Printf("Converted amount: %.2f\n", convertAmount)

	bmi_calculator.BMICalculator()
}
