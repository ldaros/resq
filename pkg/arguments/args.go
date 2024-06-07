package arguments

import (
	"flag"
	"fmt"
	"os"
)

type CommandLineArguments struct {
	SchemaFile    string
	OutputFile    string
	NumberOfLines int
}

func ParseArguments() CommandLineArguments {
	var (
		schemaFile    string
		outputFile    string
		numberOfLines int = 30
	)

	flag.StringVar(&schemaFile, "f", "", "Path to request definition file")
	flag.StringVar(&outputFile, "o", "", "Output file for the response")
	flag.IntVar(&numberOfLines, "l", 30, "Number of lines to print for JSON response")

	flag.Parse()

	if schemaFile == "" {
		fmt.Println("Please provide the request schema file using the -f flag")
		os.Exit(1)
	}

	return CommandLineArguments{SchemaFile: schemaFile, OutputFile: outputFile, NumberOfLines: numberOfLines}
}
