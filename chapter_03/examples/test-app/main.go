package main

import (
    "fmt"
    "reflect"

    "test-app/libs/p0"
    "test-app/utils"
)

func main() {
    fmt.Println("Hello", p0.Xello())

    intArr := []int{2, 3, 5, 7, 11}

    fmt.Println(reflect.TypeOf(intArr))
    fmt.Println("\n\n", p0.Name)

    p0.UseFunc(p0.SumVals, 12, 21)

    // start the casino game
    utils.TaxThePoorInMath()

}

/*
% go mod init test-app
% go run main.go

🎲 Welcome to Table 1: Lucky Draw!
You drew 98! You win the jackpot! 🏆

🎰 Welcome to Table 2: Slot Machine!
Reel 1: Lemon
Reel 2: Bell
Reel 3: Lemon
Spin complete! Did you win? Check for matching symbols!

🎡 Welcome to Table 3: Roulette!
You bet on 15. The wheel landed on 18.
💔 You lose this round. Better luck next time!

🏛️ Welcome to the High Rollers Lounge: All-In-One Game!

Turn 1:
You rolled a 6.
Lucky roll! You earn 2 points.
Mini-game: Spin the slot machine for a bonus!
Reel 1: Star
Reel 2: Diamond
Reel 3: Bell

Turn 2:
You rolled a 3.
Unlucky roll! No points this time.
Mini-game: Spin the slot machine for a bonus!
Reel 1: Bell
Reel 2: Bell
Reel 3: Lemon

Turn 3:
You rolled a 6.
Lucky roll! You earn 2 points.
Mini-game: Spin the slot machine for a bonus!
Reel 1: Lemon
Reel 2: Lemon
Reel 3: Star

Turn 4:
You rolled a 2.
Unlucky roll! No points this time.
Mini-game: Spin the slot machine for a bonus!
Reel 1: Diamond
Reel 2: Lemon
Reel 3: Cherry

Turn 5:
You rolled a 6.
Lucky roll! You earn 2 points.
Mini-game: Spin the slot machine for a bonus!
Reel 1: Cherry
Reel 2: Lemon
Reel 3: Bell

Game Over! Your final score in the High Rollers Lounge is: 6
*/
