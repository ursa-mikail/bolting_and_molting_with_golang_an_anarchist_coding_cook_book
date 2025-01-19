package main

import (
    "fmt"
    "math/rand"
    "time"

    "test-app/utils"
)

func main() {
    // Seed random number generator
    rand.Seed(time.Now().UnixNano())

    // === Launch Control: Preflight Check ===
    fmt.Println("🚀 Launch Control: Preflight Check")
    if err := utils.PreflightCheck(); err != nil {
        fmt.Printf("Mission aborted: %s\n", err)
        return
    }
    fmt.Println("All systems go! Ignition sequence start.\n")

    // === Orbit Adjustment: Error Retry ===
    fmt.Println("🛰️ Orbit Adjustment: Aligning Trajectory")
    maxRetries := 3
    for i := 1; i <= maxRetries; i++ {
        err := utils.AdjustOrbit()
        if err == nil {
            fmt.Println("Orbit adjustment successful!")
            break
        } else {
            fmt.Printf("Attempt %d failed: %s\n", i, err)
            if i == maxRetries {
                fmt.Println("Mission failed: Unable to adjust orbit.")
                return
            }
        }
    }

    // === Docking Control: Managing Multiple Errors ===
    fmt.Println("\n🤝 Docking Control: Preparing for ISS Docking")
    dockingOutcome := rand.Intn(3) // Random outcome
    switch dockingOutcome {
    case 0:
        utils.HandleDockingError(utils.ErrFuelLeak)
    case 1:
        utils.HandleDockingError(utils.ErrNavigation)
    case 2:
        fmt.Println("Docking successful! Crew transfer initiated.")
    default:
        utils.HandleDockingError(utils.ErrDockingFailed)
    }

    // === Mission Logs: Panic and Recovery ===
    fmt.Println("\n📜 Mission Logs: Embrace the Chaos")
    logs := []string{"Telemetry data received", "Solar panel deployed", "Hull breach detected", "Power system failure"}
    defer fmt.Println("Mission log complete.")
    for _, log := range logs {
        utils.ProcessLog(log)
    }

    fmt.Println("\n🎉 Mission accomplished! Returning to Earth.")
}

/*
% go mod init test-app
% go run main.go

🚀 Launch Control: Preflight Check
All systems go! Ignition sequence start.

🛰️ Orbit Adjustment: Aligning Trajectory
Orbit adjustment successful!

🤝 Docking Control: Preparing for ISS Docking
❌ Error: Navigation system failure! Switching to manual controls.

📜 Mission Logs: Embrace the Chaos
Processing log: Telemetry data received
Processing log: Solar panel deployed
Processing log: Hull breach detected
⚠️ Recovered from catastrophic error: 🚨 Critical failure: Hull integrity compromised!
Processing log: Power system failure

🎉 Mission accomplished! Returning to Earth.
Mission log complete.
*/
