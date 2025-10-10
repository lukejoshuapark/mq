package text

import (
	"fmt"
	"os"
	"strings"

	"github.com/lukejoshuapark/mq/cmd/mq/text/generate"
	"github.com/lukejoshuapark/mq/cmd/mq/text/parse"
)

func ProcessFile(inputFileName string) error {
	// Validate input file name is provided
	if inputFileName == "" {
		return fmt.Errorf("input file name is required")
	}

	// Check if input file exists
	if _, err := os.Stat(inputFileName); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", inputFileName)
	}

	file, err := parse.ParseFile(inputFileName)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	// Validate that at least one interface was found
	if len(file.Interfaces) == 0 {
		return fmt.Errorf("no interfaces found in file: %s", inputFileName)
	}

	outputFileName := inputFileName
	if strings.HasSuffix(outputFileName, ".go") {
		outputFileName = strings.TrimSuffix(outputFileName, ".go") + "_mq.go"
	} else {
		outputFileName += "_mq.go"
	}

	return generate.Generate(file, outputFileName)
}
