package main

import (
    "fmt"
    "reflect"

    "test-app/libs/p0"
    "test-app/utils"
)

func main() {
    fmt.Print(utils.Hello() + "\n")
    fmt.Print("hello 9\n")

    fmt.Println("Hello", p0.Xello())

    intArr := []int{2, 3, 5, 7, 11}

    fmt.Println(reflect.TypeOf(intArr))
    fmt.Println("\n\n", p0.Name)

    p0.UseFunc(p0.SumVals, 12, 21)

    // Example stock data
    // Array of stock data
    stocks := []struct {
        initialPrice float64
        currentPrice float64
        funds        float64
        symbol       string
    }{
        {1.93, 3.11, 5000.00, "XPP"},       // Example stock 1
        {100.00, 95.00, 2000.00, "TSLA"},   // Example stock 2
        {50.00, 60.00, 500.00, "MSFT"},     // Example stock 3
    }

    var totalFunds float64 // Variable to store the total funds
    var totalGainsArray []float64 // Slice to store gains or losses for all stocks

    // Loop through each stock and calculate properties
    for _, stock := range stocks {
        totalGains, gainLoss := utils.StockProperties(stock.initialPrice, stock.currentPrice, stock.funds, stock.symbol)
        totalGainsArray = append(totalGainsArray, totalGains) // Store total gains in the slice
        totalFunds += stock.funds // Accumulate the total funds

        // Print the results for each stock
        fmt.Printf("Stock Symbol: %s\n", stock.symbol)
        fmt.Printf("Initial Price: %.2f\n", stock.initialPrice)
        fmt.Printf("Current Price: %.2f\n", stock.currentPrice)
        fmt.Printf("Funds Invested: %.2f\n", stock.funds)
        fmt.Printf("Total Gains: %.2f\n", totalGains)
        if gainLoss {
            fmt.Println("Status: Gain\n")
        } else {
            fmt.Println("Status: Loss\n")
        }
    }

    // Calculate and print the total gains or losses
    totalGainsLoss := utils.Sum_GainsLoss(totalGainsArray...)
    fmt.Printf("Overall Total Gains/Loss: %.2f\n", totalGainsLoss)

    // Calculate interest and future value
    interest, futureValue := utils.CalculateInterest(totalFunds, 0.03, 0.5) // 3% interest for 0.5 years
    fmt.Printf("Total Current Funds: $%.2f, Interest: $%.2f, Future Value: $%.2f\n", totalFunds, interest, futureValue)

}

/*
% go mod init test-app
% go run main.go

Stock Symbol: XPP
Initial Price: 1.93
Current Price: 3.11
Funds Invested: 5000.00
Total Gains: 3056.99
Status: Gain

Stock Symbol: TSLA
Initial Price: 100.00
Current Price: 95.00
Funds Invested: 2000.00
Total Gains: -100.00
Status: Loss

Stock Symbol: MSFT
Initial Price: 50.00
Current Price: 60.00
Funds Invested: 500.00
Total Gains: 100.00
Status: Gain

Overall Total Gains/Loss: 3056.99
Total Current Funds: $7500.00, Interest: $1.12, Future Value: $7501.12
*/
