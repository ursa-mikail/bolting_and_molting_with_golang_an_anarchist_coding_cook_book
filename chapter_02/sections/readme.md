# Chapter 2: Variables and Data Types
## “Commitment Issues: Variables and Data Types”

Every language promises type safety or flexibility. `Go` tries to do both—and ends up ghosting you halfway.

You just wanted to store a value. Suddenly you’re entangled in a relationship with :=, var, type inference, type conversion, and explicit declarations like you're speed dating in a tax office.

You ask `Go`, “What are we?”

`Go` replies ... "Bye ..."

## 💢 Type Safety or Type Sadism?
You think declaring var x int = 5 means something. Then Go introduces interface{}—its black hole of desperation when nothing else works.

“Oh, you need polymorphism? Sure, just wrap everything in an empty coffin and pray you remember what’s inside.”

You ask Go, “Is it a string, a float, or a complex128?”
Go: “Figure it out with a type switch.”

## 🔥 Infer This!
`Go` says, “Look how simple this is!”

> x := 42 
> To be fair `:=` is `define` in formal syntax, and it is more appropriate than `equal to`. This is a better syntax. 

Then you try to make it behave in a real system and suddenly you’re stuck massaging types like a code chiropractor just to make the compiler shut up.

Explicit when it should infer. Inferred when you wish it warned you.
It’s like dating someone who’s either clingy or cold, but never just clear.

## 💔 Conversion Therapy
Need to convert between `int32` and `int64`? 
"You will say it explicitly. Every time. Even if you're converting 1 to 1."

Meanwhile, `float64` and `int` eye each other like two gang members in rival territories—one wrong move, and your program just dies.

## 🧟‍♂️ The Ghosts of Type Systems Past
Where are the enums? The generics (until recently)? The proper unions?
Instead, Go offers:
- Type aliases nobody uses
- Named types that confuse more than they clarify
- Interfaces that can be anything… and thus, are nothing.

It’s like Go wants you to believe in strong typing, but only in the abstract—like a religion without rituals.

## 🧨 “var”, “const”, “:=”, “type”
Too many ways to say “I want a name for this thing.”
None of them satisfying. All of them conditional.

You just wanted a single, coherent way to declare a variable.
Instead, Go hands you a buffet of awkward choices and tells you:
“Pick one. It might hurt later.”

## 👻 The Real Problem
Golang is commitment-phobic when it comes to types. It wants to be safe, but not too safe. Flexible, but not too flexible.
It’s like the language equivalent of someone who wants an open relationship, but gets jealous when you use reflection.

And don’t even mention reflection. That’s Go’s version of looking through your partner’s diary while pretending to trust them.

## 🧭 What It Should’ve Been
- A minimal, consistent, coherent type system.
- Fewer ways to say the same thing.
- More expressive tools, less compiler nagging.
- One path to variable declaration. Not a choose-your-own-misery adventure.

## Holy Matrimony Of The Types

To ensure variables are bounded in their form and format in programming—especially when juggling types like int32, int64, float64, int, uint, etc.—you need tight discipline and design enforcement.

### 1. KNOW YOUR TYPE (KYT) !! : Explicit Typing Everywhere (No Inference Games)
Don’t play coy with `:=`. Be brutally clear:

```
var count int32 = 0  // Don't let it float or mutate later
```

If you know what type you want—declare it. No guessing games. You’re writing code, not Tinder bios.

### 2. Static Analysis Tools (Gatekeepers, Not Suggestions)
Use linters like:
- golangci-lint
- staticcheck
- vet

…to ban loose typing and enforce strict usage patterns. Configure your linter like it’s your legal department: no nonsense, no exceptions.

### 3. Strong Typedefs for Meaningful Domains
Stop reusing primitive types for different purposes:
```
type UserID int64
type OrderID int64
```

They’re the same under the hood, but not semantically. This makes it impossible to mix them up without a cast—exactly the kind of friction you want.

### 4. Checking At The Receiption: Encapsulate Format Validation in Constructors
Never let variables be born without supervision:

```
func NewPercentage(v float64) (Percentage, error) {
    if v < 0.0 || v > 100.0 {
        return 0, fmt.Errorf("percentage out of bounds")
    }
    return Percentage(v), nil
}
```

No raw assignments. All variable births must go through the hospital, not the alley.

### 5. Restrict with Custom Types + Interfaces
Create bounded abstractions. If something can only be uint8, lock it behind a type:

```
type ByteValue uint8
```

Add Marshal, Unmarshal, String() to ensure it never escapes its cage.

### 6. Schema + Domain Contracts (Across the Codebase)
Document and enforce domain-wide contracts:
- What range is acceptable?
- What precision is allowed?
- What conversions are illegal?

Turn them into CI-enforced unit/integration tests, not tribal knowledge.

7. Refuse Silent Coercion
NEVER allow implicit widening or narrowing:

```
// Bad
var x int32 = 5
var y int64 = x // silent promotion

// Better
var y int64 = int64(x) // explicit, ugly, and intended
```

If the type changes, it should be loud, painful, and traceable.
Scream Before It Hurts (SBIH).

### 8. Avoid Defaulting to int or float64
These are lazy defaults. If you don’t know what width or signedness you need, you’re flying blind. Pick the tightest, most semantically meaningful type and commit.

### 9. CI Tests for Type Boundaries
Write unit tests just for boundary behavior:

```
func TestUserIDMax(t *testing.T) {
    var maxID UserID = UserID(math.MaxInt64)
    assert.NotPanics(t, func() { _ = ProcessUserID(maxID) })
}
```

Set boundaries for your Types.

### 10. Refactor Loosely-Typed Legacy Code
Treat interface{} like asbestos: legacy, dangerous, and to be removed wherever possible. Replace with generics or tagged unions (when the language supports them).

You're in a codependently typed relationship.
And you can never tell if it's type-safe or just emotionally unavailable.

### Final Rule
Treat types like contracts.
Types are promises. And every time you weaken that promise, you break the system a little more.

Stop being clever. Be deliberate.
Bound your variables like they’re volatile chemicals—not Post-It notes.

<hr>

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

# CI integration setup

CI integration setup that enforces strict typing and variable bounds using Go, combining linting, static analysis, and custom boundary tests. This assumes you're using GitHub Actions (common, extendable), but other CI tools like GitLab CI or Jenkins follow similar patterns.

### ⚙️ Step 1: golangci-lint Config (Root-level .golangci.yml)

```
run:
  timeout: 5m

linters:
  enable:
    - govet
    - staticcheck
    - errcheck
    - typecheck
    - gofmt
    - ineffassign
    - revive
    - unparam

linters-settings:
  revive:
    rules:
      - name: var-naming
        arguments:
          allow-leading-underscore: false
  staticcheck:
    checks: ["all"]
  gofmt:
    simplify: true
```

💣 Strict Typing Rule of Thumb: Ban use of `interface{}`, flag `int/float64` usage where inappropriate, and enforce custom domain types.

### 📁 Step 2: Example Go Boundary Test File
File: types/percentage_test.go

```
package types

import (
    "math"
    "testing"
)

func TestPercentageOutOfBounds(t *testing.T) {
    _, err := NewPercentage(-5)
    if err == nil {
        t.Fatal("expected error for value below 0")
    }
    _, err = NewPercentage(101)
    if err == nil {
        t.Fatal("expected error for value above 100")
    }
}

func TestPercentageValid(t *testing.T) {
    valid, err := NewPercentage(45.5)
    if err != nil {
        t.Fatal("unexpected error for valid percentage")
    }
    if valid != Percentage(45.5) {
        t.Fatal("invalid assignment")
    }
}

```

### 🚦 Step 3: GitHub Actions CI Workflow
File: .github/workflows/strict-ci.yml

```
name: Strict Type Check & Boundary Tests

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - name: Checkout Repo
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: 1.22

      - name: Install golangci-lint
        run: |
          curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.56.2

      - name: Run Linters
        run: golangci-lint run --timeout 3m

      - name: Run Unit Tests
        run: go test ./... -v

```

🧠 Add a Linter Rule to Ban Ambiguous Types
For added paranoia, add a custom check (or a script using grep in CI) to reject use of raw int, float64, or interface{}:

scripts/check-types.sh

```
#!/bin/bash

echo "Checking for disallowed types..."

if grep -r -E '\b(int|float64|interface{})\b' --exclude-dir=.git --exclude-dir=vendor --include=*.go .; then
  echo "❌ Disallowed type(s) found. Use explicitly bounded types or custom type aliases."
  exit 1
fi

echo "✅ Type check passed."

```

Add it to the GitHub Actions job:

```
      - name: Enforce Type Boundaries
        run: bash scripts/check-types.sh
```

## Extended setup that includes:
- Go generics with type bounds for controlled inputs.
- JSON Schema validation using gojsonschema (Refer: https://github.com/xeipuuv/gojsonschema).
- CI integration to test boundary violations and schema mismatches.

### 🧠 1. Go Generics with Type Bounds
Let’s define a bounded generic type for numeric constraints:

```
// pkg/constraints/constraints.go
package constraints

type BoundedNumber interface {
    ~int | ~int32 | ~int64 | ~float32 | ~float64
}

func Clamp[T BoundedNumber](value, min, max T) T {
    if value < min {
        return min
    }
    if value > max {
        return max
    }
    return value
}

```

### 🧪 2. Custom Type Example

```
// pkg/types/percentage.go
package types

import "fmt"

type Percentage float64

func NewPercentage(value float64) (Percentage, error) {
    if value < 0 || value > 100 {
        return 0, fmt.Errorf("percentage must be between 0 and 100")
    }
    return Percentage(value), nil
}

```

### 📜 3. JSON Schema for Input Validation
Install:
```
go get github.com/xeipuuv/gojsonschema

```

```
// schema/percentage.schema.json
{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "title": "PercentageInput",
  "type": "object",
  "properties": {
    "value": {
      "type": "number",
      "minimum": 0,
      "maximum": 100
    }
  },
  "required": ["value"]
}

```

### 🔍 4. Validate Inputs Against Schema

```
// pkg/validate/validate.go
package validate

import (
    "github.com/xeipuuv/gojsonschema"
)

func ValidateAgainstSchema(schemaPath, jsonData string) error {
    schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
    documentLoader := gojsonschema.NewStringLoader(jsonData)

    result, err := gojsonschema.Validate(schemaLoader, documentLoader)
    if err != nil {
        return err
    }

    if !result.Valid() {
        for _, err := range result.Errors() {
            return err
        }
    }
    return nil
}


```

### ✅ 5. Unit Test for Schema + Bounds


```
// pkg/types/percentage_test.go
package types

import (
    "testing"
    "yourapp/pkg/validate"
)

func TestJSONValidation(t *testing.T) {
    jsonInput := `{"value": 105}` // Should fail
    err := validate.ValidateAgainstSchema("./schema/percentage.schema.json", jsonInput)
    if err == nil {
        t.Fatal("expected JSON schema validation to fail")
    }
}


```
### 🏗️ 6. CI Update for Schema + Generic Bounds
Update GitHub Actions (strict-ci.yml)
```
      - name: Run Schema & Generic Bound Tests
        run: go test ./pkg/... -v

```
Add schema validation to type-bound test coverage:

```
scripts/check-schema.sh

```

```
#!/bin/bash
echo "Validating example input against schema..."
go run ./cmd/schema_validate/main.go || exit 1

```
### 🧱 7. (Optional) Type Aliases With JSON Support

```
// pkg/types/aliases.go
type Age uint8  // bounded naturally [0–255]
type Score float32 // maybe clamp [0.0 – 1.0] with Clamp[T]


```

### 🔧 JSON + Go Types Interop
If you're marshaling/unmarshaling structs with typed fields like Percentage, define MarshalJSON and UnmarshalJSON methods to enforce boundaries even if schema passes.

```
func (p *Percentage) UnmarshalJSON(b []byte) error {
    var v float64
    if err := json.Unmarshal(b, &v); err != nil {
        return err
    }
    if v < 0 || v > 100 {
        return fmt.Errorf("percentage out of bounds")
    }
    *p = Percentage(v)
    return nil
}

```

## 🧬 Summary

| Feature           | Description                                                             |
| ----------------- | ----------------------------------------------------------------------- |
| **Generics**      | Constrained `Clamp[T BoundedNumber]` to limit input ranges              |
| **JSON Schema**   | Validates payloads before even hitting logic                            |
| **CI**            | Tests for schema violations, boundary failures                          |
| **Safety**        | JSON marshalling guards to catch runtime violations                     |
| **Extensibility** | Add schemas per type (e.g., Age, Score), and plug in custom constraints |



✅ Result
- Linter-enforced formatting, safety, and unused variables.
- Custom tests for boundary violations.
- CI that catches type misuses before they touch main.
