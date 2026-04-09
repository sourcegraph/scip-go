package index_test

import (
	"flag"
	"fmt"
	"path/filepath"
	"sort"
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

			// Filter out documents outside of current directory
			var filteredDocs []*scip.Document
			for _, doc := range scipIndex.Documents {
				if strings.HasPrefix(doc.RelativePath, "..") {
					continue
				}
				if *filter != "" && !strings.Contains(doc.RelativePath, *filter) {
					continue
				}
				filteredDocs = append(filteredDocs, doc)
			}
			scipIndex.Documents = filteredDocs

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

			sourceFiles, err := testutil.FormatSnapshots(
				&scipIndex, "//", symbolFormatter, "")
			if err != nil {
				t.Fatal(err)
			}

			if len(scipIndex.ExternalSymbols) > 0 {
				sourceFiles = append(sourceFiles, scip.NewSourceFile(
					"external_symbols.txt",
					"external_symbols.txt",
					formatExternalSymbols(scipIndex.ExternalSymbols, symbolFormatter),
				))
			}

			return sourceFiles
		},
	)
}

func formatExternalSymbols(symbols []*scip.SymbolInformation, formatter scip.SymbolFormatter) string {
	var b strings.Builder

	sort.Slice(symbols, func(i, j int) bool {
		return symbols[i].Symbol < symbols[j].Symbol
	})

	for i, sym := range symbols {
		if i > 0 {
			b.WriteString("\n")
		}
		formatted, err := formatter.Format(sym.Symbol)
		if err != nil {
			formatted = sym.Symbol
		}
		b.WriteString(formatted)
		b.WriteString("\n")

		rels := make([]*scip.Relationship, len(sym.Relationships))
		copy(rels, sym.Relationships)
		sort.Slice(rels, func(i, j int) bool {
			return rels[i].Symbol < rels[j].Symbol
		})

		for _, rel := range rels {
			relFormatted, err := formatter.Format(rel.Symbol)
			if err != nil {
				relFormatted = rel.Symbol
			}
			kind := "reference"
			if rel.IsImplementation {
				kind = "implementation"
			}
			if rel.IsReference {
				kind = "reference"
			}
			if rel.IsTypeDefinition {
				kind = "type_definition"
			}
			fmt.Fprintf(&b, "  relationship %s %s\n", relFormatted, kind)
		}
	}

	return b.String()
}

// getTestdataRoot returns the absolute path to the testdata directory of this repository.
func getTestdataRoot(t *testing.T) string {
	testdata, err := filepath.Abs("../testdata")
	if err != nil {
		t.Fatalf("unexpected error getting absolute directory: %s", err)
	}

	return testdata
}
