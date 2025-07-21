package main

import (
	"fmt"
	"github.com/vladislav-gorovenko/review-golang/converter"
)

func main() {
	var usdInRub float64
	converter.USDToRUB(&usdInRub, 100)
	fmt.Println(usdInRub)
}
