package converter

func USDToRUB(usdInRub *float64, usd float64) {
	USDtoEUR := 0.99
	EURtoRUB := 100.0
	USDtoRUB := USDtoEUR * EURtoRUB
	*usdInRub = usd * USDtoRUB
}
