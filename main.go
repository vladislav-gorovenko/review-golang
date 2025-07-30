package main

import (
	"fmt"
	"github.com/vladislav-gorovenko/review-golang/converter"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	converter.Convert()

	//bmi_calculator.BMICalculator()
}
