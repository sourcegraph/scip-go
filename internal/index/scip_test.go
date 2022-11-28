package index_test

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/index"
	"github.com/sourcegraph/scip/bindings/go/scip"
	"github.com/sourcegraph/scip/bindings/go/scip/testutil"
)

// Use "update-snapshots" to update snapshots
var filter = flag.String("filter", "", "filenames to filter by")

func TestSnapshots(t *testing.T) {
	snapshotRoot := getRepositoryRoot(t)

	testutil.SnapshotTest(t,
		snapshotRoot,
		func(inputDirectory, outputDirectory string, sources []*scip.SourceFile) []*scip.SourceFile {
			if *filter != "" && !strings.Contains(inputDirectory, *filter) {
				return []*scip.SourceFile{}
			}

			index, err := index.Index(config.IndexOpts{
				ModuleRoot:    inputDirectory,
				ModuleVersion: "0.1-test",
				ModulePath:    "sg/" + filepath.Base(inputDirectory),
			})

			if err != nil {
				t.Fatal(err)
			}

			symbolFormatter := scip.DescriptorOnlyFormatter
			symbolFormatter.IncludePackageName = func(name string) bool { return !strings.HasPrefix(name, "sg/") }

			sourceFiles := []*scip.SourceFile{}
			for _, doc := range index.Documents {
				if *filter != "" && !strings.Contains(doc.RelativePath, *filter) {
					continue
				}

				formatted, err := testutil.FormatSnapshot(doc, index, "//", symbolFormatter)
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

// getRepositoryRoot returns the absolute path to the testdata directory of this repository.
func getRepositoryRoot(t *testing.T) string {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("unexpected error getting working directory: %s", err)
	}

	root, err := filepath.Abs(filepath.Join(wd, "../../internal/testdata"))
	if err != nil {
		t.Fatalf("unexpected error getting absolute directory: %s", err)
	}

	return root
}
