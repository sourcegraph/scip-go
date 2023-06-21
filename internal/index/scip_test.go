package index_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sourcegraph/scip-go/internal/command"
	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/index"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"github.com/sourcegraph/scip/bindings/go/scip/testutil"
	"google.golang.org/protobuf/proto"
)

// Use "update-snapshots" to update snapshots
var filter = flag.String("filter", "", "filenames to filter by")

func TestSnapshots(t *testing.T) {
	snapshotRoot := getTestdataRoot(t)

	testutil.SnapshotTest(t,
		snapshotRoot,
		func(inputDirectory, outputDirectory string, sources []*scip.SourceFile) []*scip.SourceFile {
			fmt.Println("Indexing", inputDirectory, "outputDirectory", outputDirectory)

			if *filter != "" && !strings.Contains(inputDirectory, *filter) {
				return []*scip.SourceFile{}
			}

			scipIndex := scip.Index{}
			writer := func(msg proto.Message) {
				switch msg := msg.(type) {
				case *scip.Metadata:
					scipIndex.Metadata = msg
				case *scip.Document:
					scipIndex.Documents = append(scipIndex.Documents, msg)
				case *scip.SymbolInformation:
					scipIndex.ExternalSymbols = append(scipIndex.ExternalSymbols, msg)
				}
			}

			err := index.Index(writer, config.IndexOpts{
				ModuleRoot:      inputDirectory,
				ModuleVersion:   "0.1.test",
				ModulePath:      "sg/" + filepath.Base(inputDirectory),
				GoStdlibVersion: "go1.19",
			})

			if err != nil {
				t.Fatal(err)
			}

			symbolFormatter := scip.SymbolFormatter{
				OnError:               func(err error) error { return err },
				IncludeScheme:         func(scheme string) bool { return scheme == "local" },
				IncludePackageManager: func(_ string) bool { return false },
				IncludePackageName:    func(name string) bool { return !strings.HasPrefix(name, "sg/") },
				IncludePackageVersion: func(_ string) bool { return true },
				IncludeDescriptor:     func(_ string) bool { return true },
			}

			sourceFiles := []*scip.SourceFile{}
			for _, doc := range scipIndex.Documents {
				// Skip files outside of current directory
				if strings.HasPrefix(doc.RelativePath, "..") {
					continue
				}

				if *filter != "" && !strings.Contains(doc.RelativePath, *filter) {
					continue
				}

				formatted, err := testutil.FormatSnapshot(doc, &scipIndex, "//", symbolFormatter)
				if err != nil {
					t.Errorf("Failed to format document: %s // %s", doc.RelativePath, err)
				}

				sourceFiles = append(sourceFiles, scip.NewSourceFile(
					doc.RelativePath,
					doc.RelativePath,
					formatted,
				))
			}

			return sourceFiles
		},
	)

	// snapshots := map[string]string{}
	// for _, doc := range index.Documents {
	// 	snapshot, err := testutil.FormatSnapshot(doc, index, "//", scip.VerboseSymbolFormatter)
	// 	if err != nil {
	// 		// t.Fatal(fmt.Sprintf("failed to process %s due to %s", doc.RelativePath, err))
	// 		fmt.Println(fmt.Sprintf("failed to process %s due to %s", doc.RelativePath, err))
	// 	}
	//
	// 	snapshots[doc.RelativePath] = snapshot
	// }
	//
	// fmt.Println(snapshotRoot)
}

// getTestdataRoot returns the absolute path to the testdata directory of this repository.
func getTestdataRoot(t *testing.T) string {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("unexpected error getting working directory: %s", err)
	}

	root, err := command.Run(wd, "git", "rev-parse", "--show-toplevel")
	if err != nil {
		t.Fatalf("unexpected error getting working directory: %s", err)
	}

	testdata, err := filepath.Abs(filepath.Join(root, "internal/testdata"))
	if err != nil {
		t.Fatalf("unexpected error getting absolute directory: %s", err)
	}

	return testdata
}
