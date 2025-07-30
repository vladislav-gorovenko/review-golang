package converter

import (
	"errors"
	"fmt"
)

type Currency string

const USDtoRUB float64 = 80

const (
	USD Currency = "USD"
	RUB Currency = "RUB"
)

func getCurrencyFrom() (Currency, error) {
	var fromCurrency Currency
	fmt.Print("Type currency from (USD, RUB): ")
	_, err := fmt.Scan(&fromCurrency)
	if err != nil {
		return "", err
	}
	return fromCurrency, nil
}

func getCurrencyTo() (Currency, error) {
	var toCurrency Currency
	fmt.Print("Type currency to (USD, RUB): ")
	_, err := fmt.Scan(&toCurrency)
	if err != nil {
		return "", err
	}
	return toCurrency, nil
}

func GetCurrencyInput() (Currency, Currency, error) {
	fromCurrency, errFrom := getCurrencyFrom()
	toCurrency, errTo := getCurrencyTo()
	if errFrom != nil || errTo != nil {
		return "", "", errors.New("Can't get currency input")
	}
	return fromCurrency, toCurrency, nil
}

func GetAmountInput() (float64, error) {
	var amount float64

	fmt.Print("Type amount: ")
	_, err := fmt.Scan(&amount)
	if err != nil {
		return 0, err
	}

	return amount, nil
}

func ConvertCurrency(from Currency, to Currency, amount *float64) {
	if from == USD && to == RUB {
		*amount = *amount * USDtoRUB
	}
	if from == RUB && to == USD {
		*amount = *amount / USDtoRUB
	}
}

func Convert() {
	var amount float64
	var errAmount error
	from, to, errCurrency := GetCurrencyInput()
	amount, errAmount = GetAmountInput()

	if errCurrency != nil || errAmount != nil {
		panic("Invalid values provided")
	}

	ConvertCurrency(from, to, &amount)

	fmt.Print(amount)
}
