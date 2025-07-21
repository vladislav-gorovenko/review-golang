package main

import (
	"fmt"
	"github.com/vladislav-gorovenko/review-golang/converter"
	"math"
)

func imt() {
	var imt float64
	fmt.Println("Calculator of IMT")
	userWeight, userHeight := getWeightAndHeight()
	calculateIMT(&imt, userWeight, userHeight)
	printFinalResult(imt)
}

func calculateIMT(imt *float64, userWeight float64, userHeight float64) {
	const IMTPower float64 = 2
	*imt = userWeight / math.Pow(userHeight, IMTPower)
}

func printFinalResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела %.2f\n", imt)
	fmt.Print(result)
}

func getWeightAndHeight() (weight float64, height float64) {
	var userWeight float64
	var userHeight float64

	fmt.Print("Type your weight: ")
	_, errWeight := fmt.Scan(&userWeight)
	if errWeight != nil {
		return
	}
	fmt.Print("Type your height: ")
	_, errHeight := fmt.Scan(&userHeight)
	if errHeight != nil {
		return
	}

	return userWeight, userHeight
}

func main() {
	var usdInRub float64
	imt()
	converter.USDToRUB(&usdInRub, 100)
	fmt.Println(usdInRub)
}
