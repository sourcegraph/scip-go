package output

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

var logLevel = new(slog.LevelVar)

// WithProgress prints the task name, runs fn, and optionally prints elapsed time.
func WithProgress(name string, fn func() error) error {
	if logLevel.Level() > slog.LevelWarn {
		return fn()
	}
	start := time.Now()
	fmt.Printf("%s\n", name)
	err := fn()
	if logLevel.Level() < slog.LevelWarn {
		fmt.Printf("Finished in %s.\n\n", HumanElapsed(start))
	}
	return err
}

// WithProgressParallel will continuously print progress to stdout until the given wait group
// counter goes to zero. Progress is determined by the values of `c` (number of tasks completed)
// and the value `n` (total number of tasks).
func WithProgressParallel(wg *sync.WaitGroup, name string, c *uint64, n uint64) {
	if logLevel.Level() > slog.LevelWarn {
		wg.Wait()
		return
	}

	start := time.Now()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	printProgress := func() {
		completed := atomic.LoadUint64(c)
		fmt.Printf("\r%s [%d/%d]", name, completed, n)
	}

	printProgress()
	for {
		select {
		case <-done:
			fmt.Printf("\r%s [%d/%d]", name, n, n)
			if logLevel.Level() < slog.LevelWarn {
				fmt.Printf(" Finished in %s.", HumanElapsed(start))
			}
			fmt.Println()
			return
		case <-ticker.C:
			printProgress()
		}
	}
}

// SetOutputOptions configures logging based on the provided slog.Level.
//
//   - Above LevelWarn (e.g. quiet mode): all output suppressed
//   - LevelWarn: default output, no timing
//   - LevelInfo: verbose output with timing
//   - LevelDebug: very verbose output with timing
func SetOutputOptions(level slog.Level) {
	logLevel.Set(level)
	if level > slog.LevelWarn {
		slog.SetDefault(slog.New(slog.DiscardHandler))
	} else {
		slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})))
	}
}
