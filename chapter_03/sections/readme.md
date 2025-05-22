# Chapter 3: Control Structures — Because You’re Not in Control
## "Control Freak: Ifs, Loops, and Other Delusions"

Control structures promise you power. They hand you if, for, and switch—like you're the conductor of logic. But here’s the crux: you're not in control. The compiler, the runtime, the scheduler—they are. You write logic, but it doesn't guarantee behavior. Not consistently. Not deterministically. Not across all inputs.

## 🧱 The If-Else Lie
You think if gives you choice. It doesn’t. It gives you a guess—on what the input might be, on what the branch should do. Add concurrency, async IO, external state, and your "if" becomes a coin toss wrapped in hope.

## 🔁 Loops: The Illusion of Progress
Endless loops? Missed break conditions? Off-by-one at either ends?
You tell the machine to loop—but you forgot to tell it why, how long, or when to die.
It will run. And run. Until you’re debugging why your memory is gone and your server died.

## 🧃 Switch: The Pretend Pattern Matcher
Go’s switch isn't exhaustive. It doesn't force you to cover all cases.
It’s not pattern matching. It’s just a cascade of ifs in a trench coat.
Miss a default? Skip a case? Watch undefined behavior emerge like rot from wood.

## 🧠 Concurrency Makes It Worse
Now throw in goroutines. Race conditions. Timing. Your tidy if is now subject to when and who got there first.
You’re writing "control" statements in a system where the order of execution is a suggestion.

## ✅ What You Should Really Be Doing
- Treat every branch as a fault line — test it.
- Avoid stateful conditions across threads — they lie.
- Instrument loops with failsafes — max count, timeout, panic-guard.
- Use exhaustive matching — even when not required. Add the default, handle the odd case.
- Understand that “control” in programming means shaping outcomes, not guaranteeing them.



<hr>

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