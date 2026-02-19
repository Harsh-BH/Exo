package exo

import (
	"fmt"
	"time"
)

// spinner runs a simple terminal spinner while work is executed.
// It stops automatically when the returned stop function is called.
//
//	stop := startSpinner("Generating Helm chart")
//	err  := doWork()
//	stop(err)
func startSpinner(label string) func(err error) {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	done := make(chan struct{})

	go func() {
		i := 0
		for {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\r  %s  %s ", frames[i%len(frames)], label)
				time.Sleep(80 * time.Millisecond)
				i++
			}
		}
	}()

	return func(err error) {
		close(done)
		// Clear the spinner line
		fmt.Printf("\r  %s  %s \n", func() string {
			if err != nil {
				return "✗"
			}
			return "✓"
		}(), label)
	}
}
