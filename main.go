package main

import (
	"fmt"
	"github.com/vladislav-gorovenko/review-golang/converter"
)

func main() {
	var convertAmount = converter.GetAmountInput()
	fromCurrency, toCurrency := converter.GetCurrencyInput()
	converter.ConvertCurrency(fromCurrency, toCurrency, &convertAmount)
	fmt.Printf("Converted amount: %.2f\n", convertAmount)
}
