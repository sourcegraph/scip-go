package loader

import (
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/charmbracelet/log"
)

var (
	stdlibOnce sync.Once
	stdlibMap  map[string]struct{}
)

func getStdlibPackages() map[string]struct{} {
	stdlibOnce.Do(func() {
		stdlibMap = stdPackages

		if os.Getenv("GOPACKAGESDRIVER") != "" {
			log.Debug("GOPACKAGESDRIVER is set, using hardcoded stdlib list")
			return
		}

		cmd, err := exec.LookPath("go")
		if err != nil {
			log.Debug("go command not found, using hardcoded stdlib list")
			return
		}

		output, err := exec.Command(cmd, "list", "std").Output()
		if err != nil {
			log.Debug("Failed to run 'go list std', using hardcoded stdlib list", "error", err)
			return
		}

		stdlibMap = make(map[string]struct{})
		packages := strings.SplitSeq(strings.TrimSpace(string(output)), "\n")
		for pkg := range packages {
			if pkg == "" {
				continue
			}
			base := strings.Split(pkg, "/")[0]
			stdlibMap[base] = struct{}{}
		}
		
		log.Debug("Successfully loaded stdlib packages dynamically", "count", len(stdlibMap))
	})
	return stdlibMap
}
