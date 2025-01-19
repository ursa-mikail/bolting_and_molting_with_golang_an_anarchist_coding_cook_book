package utils

import (
    "errors"
    "fmt"
    "math/rand"
)

// Custom error types
var (
    ErrFuelLeak      = errors.New("critical fuel leak detected")
    ErrNavigation    = errors.New("navigation system failure")
    ErrDockingFailed = errors.New("docking procedure failed")
)

// === Preflight Check ===
func PreflightCheck() error {
    // Simulate a random error during preflight check
    if rand.Float32() < 0.2 {
        return errors.New("preflight sensor malfunction")
    }
    return nil
}

// === Orbit Adjustment ===
func AdjustOrbit() error {
    // Simulate a 50% chance of failure
    if rand.Float32() < 0.5 {
        return errors.New("thruster misalignment detected")
    }
    return nil
}

// === Handle Docking Errors ===
func HandleDockingError(err error) {
    switch err {
    case ErrFuelLeak:
        fmt.Println("❌ Critical error: Fuel leak detected! Abort docking.")
    case ErrNavigation:
        fmt.Println("❌ Error: Navigation system failure! Switching to manual controls.")
    case ErrDockingFailed:
        fmt.Println("❌ Error: Docking procedure failed. Attempting emergency protocol.")
    default:
        fmt.Println("❌ Unknown error occurred during docking.")
    }
}

// === Mission Logs: Panic and Recovery ===
func ProcessLog(log string) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Printf("⚠️ Recovered from catastrophic error: %v\n", r)
        }
    }()
    fmt.Printf("Processing log: %s\n", log)

    // Simulate a critical failure
    if log == "Hull breach detected" {
        panic("🚨 Critical failure: Hull integrity compromised!")
    }
}


