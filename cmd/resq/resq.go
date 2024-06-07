package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"resq/pkg/arguments"
	"resq/pkg/color"
	http_utils "resq/pkg/http"
	"resq/pkg/request"
	"resq/pkg/utils"
	"strings"
	"time"
)

func main() {
	args := arguments.ParseArguments()

	data := readFile(args.SchemaFile)
	schema := parseRequestSchema(data)
	req := createRequest(schema)

	printRequest(req, args)

	start := time.Now()
	resp := executeRequest(req)
	elapsed := time.Since(start)

	printResponse(args, resp, elapsed)
}

func readFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)
	utils.CheckError(err, "reading request schema file")
	return data
}

func parseRequestSchema(data []byte) request.RequestSchema {
	schema, err := request.ParseRequestSchema(data)
	utils.CheckError(err, "parsing request schema")
	return schema
}

func createRequest(schema request.RequestSchema) *http.Request {
	req, err := request.CreateRequest(schema)
	utils.CheckError(err, "creating request")
	return req
}

func executeRequest(req *http.Request) *http.Response {
	resp, err := request.ExecuteRequest(req)
	utils.CheckError(err, "executing request")
	return resp
}

func printRequest(req *http.Request, args arguments.CommandLineArguments) {
	if args.ShowRequestURL {
		fmt.Printf("Request URL:\n%s\n\n", req.URL.String())
	}

	if args.ShowRequestBody && req.Body != nil {
		fmt.Println("Request body:")
		body, err := io.ReadAll(req.Body)
		utils.CheckError(err, "reading request body")
		fmt.Printf("%s\n\n", string(body))
		req.Body = io.NopCloser(strings.NewReader(string(body)))
	}

	if args.ShowRequestHeaders {
		fmt.Println("Request headers:")
		for name, values := range req.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}
		fmt.Println()
	}
}

func printResponse(args arguments.CommandLineArguments, resp *http.Response, elapsed time.Duration) {
	defer resp.Body.Close()

	if args.ShowResponseHeaders {
		fmt.Println("Response headers:")
		for name, values := range resp.Header {
			for _, value := range values {
				fmt.Printf("%s: %s\n", name, value)
			}
		}
		fmt.Println()
	}

	body := readResponseBody(resp)

	statusStr, elapsedStr := formatResponseDetails(resp.StatusCode, elapsed)

	output := formatOutput(resp, body)

	if args.OutputFile != "" {
		writeToFile(args.OutputFile, output)
	} else {
		fmt.Println(string(utils.LimitLines(output, args.NumberOfLines)))
		fmt.Printf("Request finished with status code %s in %s\n",
			color.Colorize(statusStr, color.GetHTTPStatusColors(http_utils.Status(resp.StatusCode).Code())),
			color.Colorize(elapsedStr, color.GetExecutionTimeColor(elapsed.Milliseconds())))
	}
}

func readResponseBody(resp *http.Response) []byte {
	body, err := io.ReadAll(resp.Body)
	utils.CheckError(err, "reading response body")
	return body
}

func formatResponseDetails(statusCode int, elapsed time.Duration) (string, string) {
	status := http_utils.Status(statusCode)
	statusStr := fmt.Sprintf("%d (%s)", statusCode, status.String())
	elapsedStr := elapsed.String()
	return statusStr, elapsedStr
}

func formatOutput(resp *http.Response, body []byte) []byte {
	var output []byte
	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		formatted, err := utils.PrettyJSON(body)
		if err != nil {
			output = body
		} else {
			output = formatted
		}
	} else {
		output = body
	}
	return output
}

func writeToFile(outputFile string, output []byte) {
	err := os.WriteFile(outputFile, output, 0644)
	utils.CheckError(err, "writing response to file")
}
