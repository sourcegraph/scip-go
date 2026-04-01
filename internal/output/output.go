package output

import (
	"fmt"
	"log/slog"
	"os"
	"sync"
	"time"

	"github.com/sourcegraph/scip-go/internal/parallel"
)

var logLevel = new(slog.LevelVar)

// showTiming controls whether elapsed time is printed for progress tasks.
var showTiming bool

// WithProgress prints a spinner while the given function is active.
func WithProgress(name string, fn func() error) error {
	ch := make(chan func() error, 1)
	ch <- fn
	close(ch)

	wg, errCh, count := parallel.Run(ch)
	WithProgressParallel(wg, name, count, 1)

	// Handle any associated errors
	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
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
	fmt.Printf("%s\n", name)
	wg.Wait()

	if showTiming {
		fmt.Printf("Finished in %s.\n\n", HumanElapsed(start))
	}
}

// SetOutputOptions configures logging based on the provided slog.Level.
//
//   - Above LevelWarn (e.g. quiet mode): all output suppressed
//   - LevelWarn: default output, no timing
//   - LevelInfo: verbose output with timing
//   - LevelDebug: very verbose output with timing
func SetOutputOptions(level slog.Level) {
	if level > slog.LevelWarn {
		slog.SetDefault(slog.New(slog.DiscardHandler))
		logLevel.Set(level)
		showTiming = false
		return
	}

	logLevel.Set(level)
	showTiming = level < slog.LevelWarn
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})))
}

// Logf is a printf-style function suitable for use as packages.Config.Logf.
// It logs at Debug level.
func Logf(format string, a ...any) {
	slog.Debug(fmt.Sprintf(format, a...))
}
