
We automate the scaffolding of the <a href="https://github.com/ursa-mikail/golang-gaia-basic-structure/tree/main"> golang-gaia-basic-structure</a>.

## Breakdown of Controls Across Tables:
<pre>
Table 1: Lucky Draw
Control: if-else
Usage: Determines the size of the prize based on a random number.

Table 2: Slot Machine
Control: for loop
Usage: Spins three reels and prints the result for each spin.

Table 3: Roulette
Control: switch
Usage: Handles different outcomes based on the bet and the roulette result.


High Rollers Lounge: All-In-One Game
Control: Combines if-else, for, and switch
Usage: Includes a dice-based mini-game, slot machine spins, and bonus events.

Enjoy the illusion of winning in the casino experience, but remember: the house always wins! 🎲🎰


Lottery: A tax on people who are bad at math. - Ambrose Bierce (June 24, 1842 to 1914?) 

The lottery is a tax on poor people and on people who can’t do math. Rich people and smart people would be in the line if the lottery were a real wealth-building tool, but the truth is that the lottery is a rip-off instituted by our government. This is not a moral position; it is a mathematical, statistical fact. Studies show that the zip codes that spend four times what anyone else does on lottery tickets are those in lower-income parts of town. The lottery, or gambling of any kind, offers false hope, not a ticket out. ― Dave Ramsey, The Total Money Makeover: A Proven Plan for Financial Fitness

</pre>

<pre>
chmod +x make_go.sh
# Run the script with your desired module name:
# ./make_go.sh example.com/demo
./make_go.sh test-app

# Resulting Structure
After running the script, the structure will look like this:

test-app # or example.com/demo
├── go.mod
├── lib
│   └── p0
│       └── p0.go
├── main.go
└── util
    └── util_00.go

# modify the generated main.go, util_00.go

# test run:
% cd test-app 
% go run main.go

out:
</pre>
```
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
```