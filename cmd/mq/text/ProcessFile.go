package text

import (
	"strings"

	"github.com/lukejoshuapark/mq/cmd/mq/text/generate"
	"github.com/lukejoshuapark/mq/cmd/mq/text/parse"
)

func ProcessFile(inputFileName string) error {
	file, err := parse.ParseFile(inputFileName)
	if err != nil {
		return err
	}

	outputFileName := inputFileName
	if strings.HasSuffix(outputFileName, ".go") {
		outputFileName = strings.TrimSuffix(outputFileName, ".go") + "_mq.go"
	} else {
		outputFileName += "_mq.go"
	}

	return generate.Generate(file, outputFileName)
}
