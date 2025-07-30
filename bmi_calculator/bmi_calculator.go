package bmi_calculator

import (
	"fmt"
	"math"
)

func BMICalculator() {
	weight := getCustomerWeight()
	height := getCustomerHeight()
	bmi := getCustomerBMI(height, weight)
	bmiDescription := getBMIDescription(bmi)
	printCustomerBMIResult(bmi, bmiDescription)
}

type BMI string

const (
	SevereThinness   BMI = "Severe Thinness"
	ModerateThinness     = "Moderate Thinness"
	MildThinness         = "Mild Thinness"
	Normal               = "Normal"
	Overweight           = "Overweight"
	ObeseClass1          = "Obese Class I"
	ObeseClass2          = "Obese Class II"
	ObeseClass3          = "Obese Class III"
)

func getBMIDescription(bmi float64) BMI {
	switch {
	case bmi < 16:
		return SevereThinness
	case bmi < 17:
		return ModerateThinness
	case bmi < 18.5:
		return MildThinness
	case bmi < 25:
		return Normal
	case bmi < 30:
		return Overweight
	case bmi < 35:
		return ObeseClass1
	case bmi < 40:
		return ObeseClass2
	default:
		return ObeseClass3
	}
}

func getCustomerHeight() float64 {
	var height float64
	fmt.Print("Enter your height: ")
	_, err := fmt.Scan(&height)
	if err != nil {
		return 0
	}
	return height
}

func getCustomerWeight() float64 {
	var weight float64
	fmt.Print("Enter your weight: ")
	_, err := fmt.Scan(&weight)
	if err != nil {
		return 0
	}
	return weight
}

func getCustomerBMI(height float64, weight float64) float64 {
	return weight / math.Pow(height, 2)
}

func printCustomerBMIResult(bmi float64, description BMI) {
	fmt.Printf("Your BMI is %v, which means - %v", bmi, description)
}
