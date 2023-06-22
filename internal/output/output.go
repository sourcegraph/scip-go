package output

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
	"time"

	"github.com/efritz/pentimento"
	"github.com/sourcegraph/scip-go/internal/parallel"
)

type Verbosity int

const (
	NoOutput Verbosity = iota
	DefaultOutput
	VerboseOutput
	VeryVerboseOutput
	VeryVeryVerboseOutput
)

type Options struct {
	Verbosity      Verbosity
	ShowAnimations bool
}

var opts Options = Options{
	Verbosity:      DefaultOutput,
	ShowAnimations: false,
}

// updateInterval is the duration between updates in withProgress.
var updateInterval = time.Second / 4

// ticker is the animated throbber used in printProgress.
var ticker = pentimento.NewAnimatedString([]string{
	"⠸", "⠼",
	"⠴", "⠦",
	"⠧", "⠇",
	"⠏", "⠋",
	"⠙", "⠹",
}, updateInterval)

// var failurePrefix = "✗"
var successPrefix = "✔"

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
	sync := make(chan struct{})
	go func() {
		wg.Wait()
		close(sync)
	}()

	withTitle(name, func(printer *pentimento.Printer) {
		for {
			select {
			case <-sync:
				return
			case <-time.After(updateInterval):
			}

			printProgress(printer, name, c, n)
		}
	})
}

// withTitle invokes withTitleAnimated withTitleStatic depending on the value of animated.
func withTitle(name string, fn func(printer *pentimento.Printer)) {
	if opts.Verbosity == NoOutput {
		fn(nil)
	} else if !opts.ShowAnimations || opts.Verbosity >= VeryVerboseOutput {
		withTitleStatic(name, opts.Verbosity, fn)
	} else {
		withTitleAnimated(name, opts.Verbosity, fn)
	}
}

// withTitleStatic invokes the given function with non-animated output.
func withTitleStatic(name string, verbosity Verbosity, fn func(printer *pentimento.Printer)) {
	start := time.Now()
	fmt.Printf("%s\n", name)
	fn(nil)

	if verbosity > DefaultOutput {
		fmt.Printf("Finished in %s.\n\n", HumanElapsed(start))
	}
}

// withTitleStatic invokes the given function with animated output.
func withTitleAnimated(name string, verbosity Verbosity, fn func(printer *pentimento.Printer)) {
	start := time.Now()
	fmt.Printf("%s %s... ", ticker, name)

	_ = pentimento.PrintProgress(func(printer *pentimento.Printer) error {
		defer func() {
			_ = printer.Reset()
		}()

		fn(printer)
		return nil
	})

	fmt.Printf("%s %s... Done (%s)\n", successPrefix, name, HumanElapsed(start))
}

// printProgress outputs a throbber, the given name, and the given number of tasks completed to
// the given printer.
func printProgress(printer *pentimento.Printer, name string, c *uint64, n uint64) {
	if printer == nil {
		return
	}

	content := pentimento.NewContent()

	if c == nil {
		content.AddLine("%s %s...", ticker, name)
	} else {
		content.AddLine("%s %s... %d/%d\n", ticker, name, atomic.LoadUint64(c), n)
	}

	printer.WriteContent(content)
}

func SetOutputOptions(verb Verbosity, animation bool) {
	opts.Verbosity = verb
	opts.ShowAnimations = animation
}

func Println(a ...any) {
	if opts.Verbosity != NoOutput {
		fmt.Println(a...)
	}
}

func Printf(format string, a ...any) {
	if opts.Verbosity >= VerboseOutput {
		log.Printf(format, a...)
	}
}

func Logf(format string, a ...any) {
	if opts.Verbosity >= VeryVeryVerboseOutput {
		log.Printf(format, a...)
	}
}
