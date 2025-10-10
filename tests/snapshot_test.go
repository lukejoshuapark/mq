package tests

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/lukejoshuapark/mq/cmd/mq/text"
)

var update = flag.Bool("update", false, "update snapshot files")

// TestCase represents a single snapshot test case
type TestCase struct {
	Name      string
	Directory string
}

// getTestCases returns all test cases in the tests directory
func getTestCases(t *testing.T) []TestCase {
	testsDir := "."
	entries, err := os.ReadDir(testsDir)
	if err != nil {
		t.Fatalf("Failed to read tests directory: %v", err)
	}

	var cases []TestCase
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		// Check if directory has input.go
		inputPath := filepath.Join(testsDir, entry.Name(), "input.go")
		if _, err := os.Stat(inputPath); err == nil {
			cases = append(cases, TestCase{
				Name:      entry.Name(),
				Directory: filepath.Join(testsDir, entry.Name()),
			})
		}
	}

	return cases
}

// TestSnapshots runs snapshot tests for all test cases
func TestSnapshots(t *testing.T) {
	cases := getTestCases(t)

	if len(cases) == 0 {
		t.Fatal("No test cases found")
	}

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			testSnapshot(t, tc)
		})
	}
}

func testSnapshot(t *testing.T, tc TestCase) {
	inputPath := filepath.Join(tc.Directory, "input.go")
	expectedPath := filepath.Join(tc.Directory, "expected.go")
	actualPath := filepath.Join(tc.Directory, "input_mq.go")

	// Clean up any existing generated file
	defer func() {
		if !*update {
			os.Remove(actualPath)
		}
	}()

	// Generate the mock
	if err := text.ProcessFile(inputPath); err != nil {
		t.Fatalf("Failed to generate mock: %v", err)
	}

	// Check that the file was generated
	if _, err := os.Stat(actualPath); os.IsNotExist(err) {
		t.Fatalf("Generated file does not exist: %s", actualPath)
	}

	// Read the generated content
	actualContent, err := os.ReadFile(actualPath)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	if *update {
		// Update mode: write the generated content to the expected file
		if err := os.WriteFile(expectedPath, actualContent, 0644); err != nil {
			t.Fatalf("Failed to update snapshot file: %v", err)
		}
		t.Logf("Updated snapshot: %s", expectedPath)

		// Also clean up the generated file
		os.Remove(actualPath)
	} else {
		// Test mode: compare the generated content with the expected content
		expectedContent, err := os.ReadFile(expectedPath)
		if err != nil {
			t.Fatalf("Failed to read expected file: %v", err)
		}

		if string(actualContent) != string(expectedContent) {
			// Write actual output to a temp file for easier debugging
			tempActual := filepath.Join(tc.Directory, "actual.go")
			os.WriteFile(tempActual, actualContent, 0644)

			t.Errorf("Generated output does not match snapshot.\n"+
				"Expected: %s\n"+
				"Actual: %s\n"+
				"Run 'go test -update' to update snapshots.",
				expectedPath, tempActual)
		}
	}
}

// TestSnapshotFilesExist verifies that all test cases have expected.go files
func TestSnapshotFilesExist(t *testing.T) {
	cases := getTestCases(t)

	for _, tc := range cases {
		expectedPath := filepath.Join(tc.Directory, "expected.go")
		if _, err := os.Stat(expectedPath); os.IsNotExist(err) {
			t.Errorf("Missing expected.go for test case: %s", tc.Name)
		}
	}
}

// ExampleUsage demonstrates how to run the tests
func ExampleUsage() {
	fmt.Println("Run tests:")
	fmt.Println("  go test ./tests")
	fmt.Println()
	fmt.Println("Update snapshots:")
	fmt.Println("  go test ./tests -update")
	fmt.Println()
	fmt.Println("View changes:")
	fmt.Println("  git diff tests/")
}
