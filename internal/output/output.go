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

type Verbosity int

const (
	NoOutput Verbosity = iota
	DefaultOutput
	VerboseOutput
	VeryVerboseOutput
	VeryVeryVerboseOutput
)

var verbosity Verbosity = DefaultOutput

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
	if verbosity == NoOutput {
		wg.Wait()
		return
	}

	start := time.Now()
	fmt.Printf("%s\n", name)
	wg.Wait()

	if verbosity > DefaultOutput {
		fmt.Printf("Finished in %s.\n\n", HumanElapsed(start))
	}
}

func SetOutputOptions(verb Verbosity) {
	var handler slog.Handler
	switch verb {
	case NoOutput:
		handler = slog.DiscardHandler
	case DefaultOutput:
		logLevel.Set(slog.LevelWarn)
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})
	case VerboseOutput:
		logLevel.Set(slog.LevelInfo)
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})
	case VeryVerboseOutput, VeryVeryVerboseOutput:
		logLevel.Set(slog.LevelDebug)
		handler = slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel})
	}
	slog.SetDefault(slog.New(handler))
	verbosity = verb
}

func Logf(format string, a ...any) {
	if verbosity >= VeryVeryVerboseOutput {
		slog.Info(fmt.Sprintf(format, a...))
	}
}
