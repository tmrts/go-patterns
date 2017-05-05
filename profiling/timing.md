# Timing Functions

When optimizing code, sometimes a quick and dirty time measurement is required
as opposed to utilizing profiler tools/frameworks to validate assumptions.

Time measurements can be performed by utilizing `time` package and `defer` statements.

Also, you can use `TimeTracker` to track each step cost for a function.

## Implementation

```go
package profile

import (
	"fmt"
	"log"
	"time"
)

func Duration(invocation time.Time, name string) {
	elapsed := time.Since(invocation)
	log.Printf("%s lasted %s", name, elapsed)
}

type timeEvent struct {
	event string
	cost  time.Duration
}

type TimeTracker struct {
	action string
	start  time.Time
	events []timeEvent
}

func NewTrack(action string) *TimeTracker {
	now := time.Now()
	return &TimeTracker{
		action: action,
		start:  now,
		events: make([]timeEvent, 0, 5),
	}
}

func (tk *TimeTracker) Track(name string) {
	curEvent := timeEvent{
		event: name,
		cost:  time.Since(tk.start),
	}
	tk.events = append(tk.events, curEvent)
	tk.start = time.Now()
}

func (tk *TimeTracker) Print() {
	total := time.Now().Sub(tk.start)
	var pairs string
	for _, ev := range tk.events {
		pairs += fmt.Sprintf("%s:%s ", ev.event, ev.cost)
	}
	if pairs != "" {
		log.Printf("timeTrack:%s total lasted:%s %s\n", tk.action, total, pairs)
	}
}
```

## Usage

```go
func BigIntFactorial(x big.Int) *big.Int {
    // Arguments to a defer statement is immediately evaluated and stored.
    // The deferred function receives the pre-evaluated values when its invoked.
    defer profile.Duration(time.Now(), "IntFactorial")

    y := big.NewInt(1)
    for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
        y.Mul(y, x)
    }

    return x.Set(y)
}

func IntFactorialTrack(x big.Int) *big.Int {
	tk := profile.NewTrack("IntFactorialTrack")
	defer tk.Print()

	y := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		y.Mul(y, x)
	}
	tk.Track("step_one")
	
	z := big.NewInt(1)
	for one := big.NewInt(1); x.Sign() > 0; x.Sub(x, one) {
		z.Mul(z, x)
	}
	tk.Track("step_two")
	
	return x.Set(y.Mul(y,z))
}
```
