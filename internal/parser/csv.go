package parser

import (
	"fmt"
	"strings"

	"github.com/KaramelBytes/docloom-cli/internal/analysis"
)

type csvParser struct{}

func (csvParser) CanParse(filename string) bool {
	name := strings.ToLower(filename)
	return strings.HasSuffix(name, ".csv") || strings.HasSuffix(name, ".tsv")
}

func (csvParser) Parse(_ []byte) (string, error) {
	// We don't use the content buffer here; ParseFile currently reads file and passes content.
	// For CSV we need the on-disk path. Refactor ParseFile to pass path in addition to content
	// would be ideal, but to keep compatibility, we re-open via a small hack: return an error
	// instructing the caller to use ParseFile for on-disk files.
	return "", fmt.Errorf("csv parser requires file path; use parser.ParseFile(path)")
}

// ParseCSVFile provides CSV parsing from an absolute file path to a compact summary.
func ParseCSVFile(path string) (string, error) {
	rep, err := analysis.AnalyzeCSV(path, analysis.DefaultOptions())
	if err != nil {
		return "", err
	}
	return rep.Markdown(), nil
}
