# Chapter 5: Error Handling — Embrace the Chaos
## "Oops! Embracing the Inevitable Errors"

Errors are the norm.
Every file read, every network call, every function—failure is not optional; it's guaranteed.
And yet... developers still treat errors like unwanted guests instead of the main character.

## 🧨 The Golang Way™: Explicit, Verbose, and Often Ignored
Go claims to handle errors better by forcing you to acknowledge them:

```
result, err := DoSomething()
if err != nil {
    return nil, err
}
```
In reality? Most devs either:
- Copy-paste the same if err != nil block like it’s a prayer scroll
- Swallow the error with _ because “it’s fine, this never fails”
- Wrap it lazily: fmt.Errorf("something failed: %v", err) and call it a day
- Lie through logs: "Error handled successfully" (???)

## 🙈 The Lies We Tell Ourselves
---

“We handle all our errors.”
→ You don’t. You check them. And then suppress, ignore, or log them without recovery.

“Errors are for ops to deal with.”
→ Until ops calls you at 3 AM and you realize “log and move on” was a death sentence.

“Error wrapping gives context.”
→ Not when 6 nested fmt.Errorf("while doing X: %w", err) lines give you a stack trace that reads like abstract poetry.

---

## ⚠️ The Broken Promises of "Graceful Failure"
What devs say:

> “We fail gracefully.”

What actually happens:
---
Retrying blindly until rate-limited

Defaulting to stale or broken data

Returning nil silently and praying nothing explodes

Crashing on edge cases you knew were possible but “meh”
--- 

💡 What You Should Be Doing
 - Centralize error paths — don’t scatter them. Build pipelines where errors are first-class data.
- Tag them — use errors.Is and errors.As properly. Not every error is just an “oops.”
- Expose severity — not all errors are equal. Some are logs. Some are alerts. Some are kill switches.
- Recover when safe — not everything has to bubble up to user-facing panic.
- Fail loud on corruption, fail soft on interruption. Know the difference.

## 🎯 The Core Truth
The only thing worse than not handling errors is pretending that you did.

### 🧱 1. Errors Break Flow, Not Just Code
Errors interrupt not just execution—but cognitive flow.
They derail context, discard state, and force jumps—with or without cleanup.
- Exceptions (in languages like Java, Python) behave like "goto" with a blindfold:
- You throw from deep inside a function
- You catch somewhere else entirely
- And you hope the in-between didn't leave a landmine behind

The result: messy recovery, orphaned memory, and half-written files.
And since cleanup paths aren’t uniformly enforced, breadcrumbs die ...

### 🧹 2. Breadcrumbs Die Young
If you can't catch it, at least trace it. Leave a trail of blood before you go down ...
Logs, traces, temporary variables, diagnostic metadata—
They live in RAM, not in exception bubbles. If they are at the RAM, they can be washed away when the power is out - and when the program is shut down and/or the RAM is flushed, flooded by berserk threads or processes. 

When an exception is raised:
- Temporary data is lost
- Logging may not flush
- Panic stack may skip over useful context
- Goroutines may vanish silently

You lose the why behind the what.
No `stacktrace` shows:
- What retry count was at the time
- What state machine branch you just crossed
- What assumptions you violated right before the failure

### 💥 3. Exception != Context
Throwing an error is easy.
Throwing useful context is rare.

Most languages just pass the error type and message. That’s it.
No upstream history, no tags, no breadcrumb trail unless you build it.

This leads to:

```
panic: index out of range [3] with length 3
```

Where? Why? What was the loop? What input triggered it?
We want to know ... Is it a murder? Or a suicide? 

### 🔕 4. Silenced by Structure
In many production systems:
- Breadcrumbs get filtered out (log level too low)
- Stacktraces are sanitized (security)
- Errors are aggregated (monitoring dashboards reduce signal)
- Exception handling is generic (try { ... } catch (Exception e) {})

In other words: the crime scene is scrubbed clean before you arrive ... and some murders are meant to look like - staged suicides. 

### 🧼 5. Errors Are Politically Dangerous
Let’s not pretend:
- Too many exceptions = “your code is fragile”
- Too much logging = “you’re spamming logs”
- Too many checks = “you’re slowing it down”
- Uncaught exception = “you missed a critical case”

So devs:
- Hide errors
- Mask them with generic messages
- Avoid throwing until it’s too late
- Assume “retry once” fixes everything
- Kill breadcrumbs before the real issue surfaces

✅ What Should Be Done
- Treat errors as first-class data, not control flow
- Always attach context—structured, typed, versioned
- Make breadcrumbs persistent, not transient (disk snapshots, log caches, ring buffers)
- Log before, not just when you fail
- Use semantic errors: know what kind of error happened, not just that one did



<hr>

We automate the scaffolding of the <a href="https://github.com/ursa-mikail/golang-gaia-basic-structure/tree/main"> golang-gaia-basic-structure</a>.

## Breakdown of Controls and Chaos
<pre>
Launch Control: Preflight Check
What Could Go Wrong: Preflight sensors could malfunction, causing an error.
Error Handling: If the error is detected, the mission is aborted to prevent catastrophic failure.

Orbit Adjustment: Error Retry
What Could Go Wrong: Thrusters could misalign during orbit adjustment, failing to correct trajectory.
Error Handling: Uses a loop to retry the adjustment up to three times. On repeated failure, the mission is terminated.

Docking Control: Managing Multiple Errors
What Could Go Wrong: Fuel leaks, navigation failures, or docking system malfunctions.
Error Handling: A switch statement is used to handle specific errors differently (e.g., abort docking, switch to manual controls).

Mission Logs: Embrace the Chaos
What Could Go Wrong: Critical errors like a hull breach or power failure could occur.
Error Handling: Simulates catastrophic errors with panic, but uses recover to handle them gracefully and prevent the program from crashing.

Key Takeaways:
Error Handling in Go: Use error for expected errors, panic for catastrophic errors, and recover to manage them gracefully.
Space Chaos Simulations: Randomness adds a realistic element of unpredictability.
Resilience: Even when things go very wrong, error handling ensures the program doesn’t completely crash.

Remember: In space, no one can hear you scream—but Go can handle your errors! 🚀✨
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
```