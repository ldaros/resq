package arguments

import (
	"flag"
	"fmt"
	"os"
)

type CommandLineArguments struct {
	SchemaFile          string
	OutputFile          string
	NumberOfLines       int
	ShowResponseHeaders bool
	ShowRequestHeaders  bool
	ShowRequestBody     bool
	ShowRequestURL      bool
}

func ParseArguments() CommandLineArguments {
	var (
		schemaFile         string
		outputFile         string
		numberOfLines      int = 30
		showRequestString  string
		showResponseString string
	)

	flag.StringVar(&schemaFile, "r", "", "Path to the request definition YAML file")
	flag.StringVar(&outputFile, "o", "", "Path to the output file where the response will be saved")
	flag.IntVar(&numberOfLines, "l", 30, "Number of lines to display from the JSON response")
	flag.StringVar(&showRequestString, "rq", "", "Specify which parts of the request to show: headers (h), body (b), URL (u)")
	flag.StringVar(&showResponseString, "rs", "", "Specify which parts of the response to show: headers (h)")

	flag.Parse()

	if schemaFile == "" {
		fmt.Println("Provide a path to the request definition YAML file using the -r flag")
		os.Exit(1)
	}

	showRequestHeaders, showRequestBody, showRequestURL := parseRequestFlags(showRequestString)
	showResponseHeaders := parseResponseFlags(showResponseString)

	return CommandLineArguments{
		SchemaFile:          schemaFile,
		OutputFile:          outputFile,
		NumberOfLines:       numberOfLines,
		ShowResponseHeaders: showResponseHeaders,
		ShowRequestHeaders:  showRequestHeaders,
		ShowRequestBody:     showRequestBody,
		ShowRequestURL:      showRequestURL,
	}
}

func parseRequestFlags(flagString string) (bool, bool, bool) {
	var headers, body, url bool

	for _, c := range flagString {
		switch c {
		case 'h':
			headers = true
		case 'b':
			body = true
		case 'u':
			url = true
		default:
			fmt.Println("Invalid argument for -rq")
			os.Exit(1)
		}
	}

	return headers, body, url
}

func parseResponseFlags(flagString string) bool {
	var headers bool

	for _, c := range flagString {
		switch c {
		case 'h':
			headers = true
		default:
			fmt.Println("Invalid argument for -rs")
			os.Exit(1)
		}
	}

	return headers
}
