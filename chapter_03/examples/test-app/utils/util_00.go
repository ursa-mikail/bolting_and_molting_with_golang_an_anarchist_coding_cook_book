package utils

import (
    "fmt"
    "math/rand"
    "time"
)

func TaxThePoorInMath() {
    // Seed the random number generator
    rand.Seed(time.Now().UnixNano())

    // === Table 1: Lucky Draw ===
    fmt.Println("🎲 Welcome to Table 1: Lucky Draw!")
    randomNumber := rand.Intn(100)
    if randomNumber < 30 {
        fmt.Printf("You drew %d! You win a small prize! 🎁\n", randomNumber)
    } else if randomNumber >= 30 && randomNumber <= 70 {
        fmt.Printf("You drew %d! You win a medium prize! 🎉\n", randomNumber)
    } else {
        fmt.Printf("You drew %d! You win the jackpot! 🏆\n", randomNumber)
    }

    // === Table 2: Slot Machine ===
    fmt.Println("\n🎰 Welcome to Table 2: Slot Machine!")
    symbols := []string{"Cherry", "Lemon", "Bell", "Star", "Diamond"}
    for i := 0; i < 3; i++ {
        spin := symbols[rand.Intn(len(symbols))]
        fmt.Printf("Reel %d: %s\n", i+1, spin)
    }
    fmt.Println("Spin complete! Did you win? Check for matching symbols!")

    // === Table 3: Roulette ===
    fmt.Println("\n🎡 Welcome to Table 3: Roulette!")
    bet := rand.Intn(37) // Random bet from 0 to 36
    rouletteNumber := rand.Intn(37)
    fmt.Printf("You bet on %d. The wheel landed on %d.\n", bet, rouletteNumber)

    switch {
    case bet == rouletteNumber:
        fmt.Println("🎉 Exact match! You win big!")
    case bet%2 == rouletteNumber%2:
        fmt.Println("😊 You win! The color matches (Red or Black).")
    default:
        fmt.Println("💔 You lose this round. Better luck next time!")
    }

    // === High Rollers Lounge: All-In-One Game ===
    fmt.Println("\n🏛️ Welcome to the High Rollers Lounge: All-In-One Game!")
    playerScore := 0

    for turn := 1; turn <= 5; turn++ {
        fmt.Printf("\nTurn %d:\n", turn)

        // Dice roll
        dice := rand.Intn(6) + 1
        fmt.Printf("You rolled a %d.\n", dice)

        // Decision-making with if-else
        if dice == 6 {
            fmt.Println("Lucky roll! You earn 2 points.")
            playerScore += 2
        } else if dice >= 4 {
            fmt.Println("Good roll! You earn 1 point.")
            playerScore++
        } else {
            fmt.Println("Unlucky roll! No points this time.")
        }

        // Bonus events using switch
        switch dice {
        case 1:
            fmt.Println("Bonus Event: You found a treasure chest!")
        case 5:
            fmt.Println("Bonus Event: You encountered a friendly NPC!")
        }

        // Slot machine roll (nested loop)
        fmt.Println("Mini-game: Spin the slot machine for a bonus!")
        for i := 0; i < 3; i++ {
            spin := symbols[rand.Intn(len(symbols))]
            fmt.Printf("Reel %d: %s\n", i+1, spin)
        }
    }

    fmt.Printf("\nGame Over! Your final score in the High Rollers Lounge is: %d\n", playerScore)
}

