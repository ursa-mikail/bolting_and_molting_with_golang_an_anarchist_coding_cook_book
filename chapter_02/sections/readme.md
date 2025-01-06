## Features:
### Introduction:
Introduces variables and data types as characters with quirky personalities.

#### Demonstration of Data Types:
Shows usage of string, int, float64, and bool.

#### Interaction Between Data Types:
Demonstrates type conversion and operations between incompatible types (e.g., int and float64).

#### SHA256 ID Generation:
Uses the GenerateSHA256IDs function to create unique IDs for variables.

#### Constants:
Adds a constant to emphasize the difference between variables and constants.

```
Variables may have commitment issues, but constants are forever.
```

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
├── libs
│   └── p0
│       └── p0.go
├── main.go
└── utils
    └── util_00.go

# modify the generated p0.go
# modify the generated main.go

# test run:
% cd test-app 
# note: if you have multiple *.go on the same folder that is working with main.go: % go run .
# here we have main.go referring to package `lib/p0` (`test-app/lib/p0`, as we define the module as `test-app`): 
% go run main.go

Note: To export a variable or make a function available for all packages in the app is to Capitalize the 1st letter of function name, e.g. p0.ExplainVariables() in main.go for `ExplainVariables()` from `lib/p0` is quite non-inituitive. 

out:
</pre>
```
% go run main.go
Hello from utils
Welcome to 'Commitment Issues: Variables and Data Types'!
Today, we'll explore Golang's quirky variable and data type relationships.
Commitment Issues: Why can't variables and types just get along?
- A string says, 'I'm here for the long haul, but only with text.'
- An int replies, 'Numbers are my thing, and I prefer no decimal drama.'
- A float counters, 'Decimals? Drama? No, it's elegance.'
- A bool smirks, 'True or false, I'm all about commitment.'

Meet our star:
- Name: Gopher (string)
- Age: 10 (int)
- Height: 1.75 (float64)
- Is Gopher cute? true (bool)

But wait, what happens if we try mixing them?
- Gopher says, 'Hey, can I combine age and height?'
Result: 11.75 (age + height)

Generating SHA256 IDs for some favorite variables:
- Gopher: 654276d49262121a990007f74bf1ae36f54b5e44425cae68d77399f5fbf25a5b
- 10: 4a44dc15364204a80fe80e9039455cc1608281820fe2b24f1e5233ade6af1dd5
- 1.75: f4881c772c8950930750e103abbe15b6720b84168921e66850d5800500ea0865
- true: b5bea41b6c623f7c09f1bf24dcae58ebab3c0cdd90ad966bc43a45b44867e12b

A constant reminder: Keep coding and stay quirky!

And remember: Variables may have commitment issues, but constants are forever.
```