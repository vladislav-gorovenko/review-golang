package converter

import "fmt"

type Currency string

const USDtoRUB float64 = 80

const (
	USD Currency = "USD"
	RUB Currency = "RUB"
)

func GetCurrencyInput() (Currency, Currency) {
	var fromCurrency Currency
	var toCurrency Currency

	fmt.Print("Type both from and to currency: ")
	n, errCurrency := fmt.Scan(&fromCurrency, &toCurrency)
	if n != 2 || errCurrency != nil {
		return "USD", "RUB"
	}

	return fromCurrency, toCurrency
}

func GetAmountInput() float64 {
	var amount float64

	fmt.Print("Type amount: ")
	_, errAmount := fmt.Scan(&amount)
	if errAmount != nil {
		return 0
	}

	return amount
}

func ConvertCurrency(from Currency, to Currency, amount *float64) {
	if from == USD && to == RUB {
		*amount = *amount * USDtoRUB
	}
	if from == RUB && to == USD {
		*amount = *amount / USDtoRUB
	}
}
