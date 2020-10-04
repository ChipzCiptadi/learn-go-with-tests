package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3
const sleep = "Sleep"
const write = "Write"

// Sleeper the sleeper interface to mimic time.Sleep()
type Sleeper interface {
	Sleep()
}

// DefaultSleeper the non-configurable sleeper
type DefaultSleeper struct{}

// Sleep for 1 second
func (s DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// ConfigurableSleeper allow sleeper to be configured by duration and DI time.Sleep method
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep for duration specified in struct
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown print 3 2 1 Go! and sleep for x time between each count
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintf(out, "%d\n", i)
	}
	sleeper.Sleep()
	fmt.Fprintf(out, finalWord)
}

func main() {
	// sleeper := &DefaultSleeper{}
	sleeper := &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
