package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// runCmd is a helper to execute the root command with args.
func runCmd(t *testing.T, args ...string) {
	t.Helper()
	// Reset sticky flags that may persist Changed state across invocations
	if f := rootCmd.Flags(); f != nil {
		if fl := f.Lookup("budget-limit"); fl != nil {
			_ = fl.Value.Set("0")
			fl.Changed = false
		}
		if fl := f.Lookup("prompt-limit"); fl != nil {
			_ = fl.Value.Set("0")
			fl.Changed = false
		}
		if fl := f.Lookup("print-prompt"); fl != nil {
			_ = fl.Value.Set("false")
			fl.Changed = false
		}
	}
	// Reset generateCmd flags as well
	if f := generateCmd.Flags(); f != nil {
		if fl := f.Lookup("budget-limit"); fl != nil {
			_ = fl.Value.Set("0")
			fl.Changed = false
		}
		if fl := f.Lookup("prompt-limit"); fl != nil {
			_ = fl.Value.Set("0")
			fl.Changed = false
		}
		if fl := f.Lookup("print-prompt"); fl != nil {
			_ = fl.Value.Set("false")
			fl.Changed = false
		}
	}
	// Reset bound variables
	genBudgetLimit = 0
	genPromptLimit = 0
	genPrintPrompt = false
	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		t.Fatalf("command %v failed: %v", args, err)
	}
}

func TestCLI_BudgetLimitBlocksGeneration(t *testing.T) {
	home := t.TempDir()
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)
	os.Setenv("HOME", home)

	// Create a larger doc to get a non-trivial token count
	docPath := filepath.Join(home, "doc.md")
	if err := os.WriteFile(docPath, []byte("Title\n\n"+strings.Repeat("content ", 3000)), 0o644); err != nil {
		t.Fatalf("write doc: %v", err)
	}

	runCmd(t, "init", "budget", "-d", "budget test")
	runCmd(t, "add", "-p", "budget", docPath)
	runCmd(t, "instruct", "-p", "budget", "Summarize")

	// Expect generate to fail due to very small budget
	rootCmd.SetArgs([]string{"generate", "-p", "budget", "--dry-run", "--budget-limit", "0.0001"})
	if err := rootCmd.Execute(); err == nil {
		t.Fatalf("expected error due to budget limit, got nil")
	}
}
func TestCLI_Init_Add_Instruct_GenerateDryRun(t *testing.T) {
	// Use a temp HOME to isolate config and projects
	home := t.TempDir()
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome)
	os.Setenv("HOME", home)

	// Create a doc file to add
	docPath := filepath.Join(home, "doc1.md")
	if err := os.WriteFile(docPath, []byte("# Title\n\nSome content."), 0o644); err != nil {
		t.Fatalf("write doc: %v", err)
	}

	// init project
	runCmd(t, "init", "itest", "-d", "integration test")
	// add doc
	runCmd(t, "add", "-p", "itest", docPath, "--desc", "first doc")
	// set instructions
	runCmd(t, "instruct", "-p", "itest", "Summarize the content")
	// generate dry-run with prompt limit for speed
	runCmd(t, "generate", "-p", "itest", "--dry-run", "--prompt-limit", "2000")
}
