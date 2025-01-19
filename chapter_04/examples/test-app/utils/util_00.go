package utils

func Hello() string {
    return "Hello from util"
}

// Multiple Return Values, Named Return Parameters Implicit
func StockProperties(initial_price float64, current_price float64, funds float64, symbol string) (total_gains float64, gain_loss bool) {
    total_units_purchased := funds / initial_price // Calculate total units purchased
    total_gains = (current_price - initial_price) * total_units_purchased
    gain_loss = total_gains > 0
    return // Implicitly returns `total_gains` and `gain_loss`
}

// Variadic function to calculate the sum of gains or losses
func Sum_GainsLoss(numbers ...float64) float64 {
    total := 0.0 // Initialize total as a float64
    for _, num := range numbers {
        total += num // Add each number to the total
    }
    return total // Return the sum
}

func CalculateInterest(principal, rate, years float64) (float64, float64) {
    interest := principal * (rate / 100) * years
    futureValue := principal + interest
    return interest, futureValue
}

