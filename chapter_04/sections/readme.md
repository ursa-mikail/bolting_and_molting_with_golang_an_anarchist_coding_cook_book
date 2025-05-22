# Chapter 4: Functions — The Buck Stops Here
## "Function Fiascos: Where All Your Problems Go to Die"

Functions are supposed to be the clean part.
You give them inputs. They give back outputs. Contained. Testable. Reusable.

## 🎭 The Lie of Simplicity
Go makes it look clean:

```
func Add(a int, b int) int {
    return a + b
}
```

But the moment real-world complexity walks in, you’re doing type juggling, interface wrapping, error plumbing, and context passing... and pretending it's still "simple".

## 🩹 Where the Problems Begin: 🎁 Parameters and Their Baggage
Functions are supposed to reduce scope. Instead, they drag in the entire application through global dependencies, closures, or hidden side effects.

- Too many inputs → refactor trap
- Too few inputs → function doesn’t do what it says
- Pointer vs value → performance bugs or silent state corruption

Functions become junk drawers—all the logic you couldn't figure out where to put, now shoved into one black box.
Too many parameters? No clear contract? Magic return types? Welcome to your own design debt factory. Now you are tied to a undocumented contract of what to give it, types, how many inputs, and what are the bounds of those inputs, etc? And you will be too afraid to look at the other side - the outputs, or if it even returns at all. The garbarge out may be the next problem for the garbarge in. In the same where those inputs may come from the same. One function just pass out and pass forward the same mess. It becomes a cycle of comedy, and it may boomerang back to the same karma, all that looks too familiar in the circus of the assembly line. 

### ❌ Return Values: The Silent Killers
Go’s error return pattern is supposed to be explicit.
In practice, it leads to:

```
if err != nil {
    // repeat 500 times across your codebase
}
```
Or worse, silently ignored returns:

```
result, _ := DoThing() // cool, no idea if it worked
```

## 🧪 Functions as Landmines
Testing functions sounds easy.
Until you need:
- Context mocks
- Interface substitutions
- Side effect tracking
- Dependency injection without a real DI framework

Then suddenly, your 5-line function needs 50 lines of setup to even exist in test. Then you are expected to introspect, and once you dig, you never leave that limbo. If you have to dig it, there will be a mess. Once you crawl out from that mass grave, you will never be the same coder again. We survived to say this ...

✅ What to Actually Do
- Limit scope: One function, one job. If it needs flags to change behavior, split it.
- Return structs or errors with context—not just values.
- Don’t hide side effects: Log them, mock them, or isolate them.
- Avoid default returns or swallowing errors: You’re not helping.
- Treat every function as if someone else will read and misuse it. Because they will.

Functions don’t clean up messes. They collect them.
They centralize complexity. And if you're not careful, they'll bury every bug you didn't know you had—until the day it explodes in prod.

<hr>

We automate the scaffolding of the <a href="https://github.com/ursa-mikail/golang-gaia-basic-structure/tree/main"> golang-gaia-basic-structure</a>.

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
Hello from util
hello 9
Hello World
[]int


 DefaultName
Result: 33

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
```