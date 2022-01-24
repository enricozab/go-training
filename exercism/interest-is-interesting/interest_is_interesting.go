// Package interest computes the interest rate your bank gives you depending on the amount of money in your account
package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	if balance < 0 {
		return float32(3.213)
	} else if balance < 1000 {
		return float32(0.5)
	} else if balance < 5000 {
		return float32(1.621)
	}
	return float32(2.475)
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	if balance < 0 {
		return balance * 0.03213
	} else if balance < 1000 {
		return balance * 0.005
	} else if balance < 5000 {
		return balance * 0.01621
	}
	return balance * 0.02475
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	if balance < 0 {
		return balance + (balance * 0.03213)
	} else if balance < 1000 {
		return balance + (balance * 0.005)
	} else if balance < 5000 {
		return balance + (balance * 0.01621)
	}
	return balance + (balance * 0.02475)
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance:
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	var years int

	for balance < targetBalance {
		if balance < 0 {
			balance += balance * 0.03213
		} else if balance < 1000 {
			balance += balance * 0.005
		} else if balance < 5000 {
			balance += balance * 0.01621
		} else {
			balance += balance * 0.02475
		}

		years++
	}

	return years
}
