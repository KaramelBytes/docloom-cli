package parser

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/KaramelBytes/docloom-cli/internal/analysis"
)

type xlsxParser struct{}

func (xlsxParser) CanParse(filename string) bool {
	return strings.HasSuffix(strings.ToLower(filename), ".xlsx")
}

func (xlsxParser) Parse(_ []byte) (string, error) {
	return "", fmt.Errorf("xlsx parser requires file path; use parser.ParseFile(path)")
}

// ParseXLSXFile analyzes the first sheet and returns a compact summary.
func ParseXLSXFile(path string, sheetName string, sheetIndex int) (string, error) {
	rep, err := analysis.AnalyzeXLSX(path, analysis.DefaultOptions(), sheetName, sheetIndex)
	if err != nil {
		return "", err
	}
	// Optionally, include sheet in name for clarity
	if rep != nil && rep.Name == filepath.Base(path) && sheetName != "" {
		rep.Name = fmt.Sprintf("%s (sheet: %s)", rep.Name, sheetName)
	}
	return rep.Markdown(), nil
}
