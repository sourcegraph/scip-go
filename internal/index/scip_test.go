package index_test

import (
	"flag"
	"fmt"
	"net/url"
	"path/filepath"
	"strings"
	"testing"

	"github.com/scip-code/scip/bindings/go/scip"
	"github.com/scip-code/scip/bindings/go/scip/testutil"
	"github.com/sourcegraph/scip-go/internal/config"
	"github.com/sourcegraph/scip-go/internal/index"
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
			writer := func(msg proto.Message) error {
				switch msg := msg.(type) {
				case *scip.Metadata:
					scipIndex.Metadata = msg
				case *scip.Document:
					scipIndex.Documents = append(scipIndex.Documents, msg)
				case *scip.SymbolInformation:
					scipIndex.ExternalSymbols = append(scipIndex.ExternalSymbols, msg)
				}
				return nil
			}

			err := index.Index(writer, config.IndexOpts{
				ModuleRoot:      inputDirectory,
				ModuleVersion:   "0.1.test",
				ModulePath:      "sg/" + filepath.Base(inputDirectory),
				GoStdlibVersion: "go1.22",
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
				IncludeRawDescriptor:  func(descriptor *scip.Descriptor) bool { return true },
				IncludeDisambiguator:  func(_ string) bool { return true },
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

				sourcePath, _ := url.JoinPath(scipIndex.Metadata.ProjectRoot, doc.RelativePath)
				sourceUrl, _ := url.Parse(sourcePath)
				formatted, err := testutil.FormatSnapshot(doc, "//", symbolFormatter, sourceUrl.Path)
				if err != nil {
					t.Errorf("Failed to format document: %s // %s", sourceUrl.Path, err)
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
}

// getTestdataRoot returns the absolute path to the testdata directory of this repository.
func getTestdataRoot(t *testing.T) string {
	testdata, err := filepath.Abs("../testdata")
	if err != nil {
		t.Fatalf("unexpected error getting absolute directory: %s", err)
	}

	return testdata
}
