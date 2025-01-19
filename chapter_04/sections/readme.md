
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