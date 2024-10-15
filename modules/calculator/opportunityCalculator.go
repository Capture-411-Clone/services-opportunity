package calculator

func CalculateOpportunityPrice(contractValueFloat64 float64) float64 {
	// $250
	// $1,000,000 & under

	// $1,000
	// $1,000,001 - $25,000,000

	// $1,500
	// $25,000,001 - $100,000,000

	// $2,000
	// $100,000,001 - $500,000,000

	// $2,500
	// $500,000,001 & Above

	var finalPrice float64

	switch {
	case contractValueFloat64 <= 1000000:
		// $250
		finalPrice = 25000
	case contractValueFloat64 <= 25000000:
		// $1,000
		finalPrice = 100000
	case contractValueFloat64 <= 100000000:
		// $1,500
		finalPrice = 150000
	case contractValueFloat64 <= 500000000:
		// $2,000
		finalPrice = 200000
	case contractValueFloat64 > 500000000:
		// $2,500
		finalPrice = 250000
	default:
		// $250
		finalPrice = 25000

	}

	return finalPrice

}
