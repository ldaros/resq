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

	data, err := os.ReadFile(args.SchemaFile)
	utils.CheckError(err, "reading request schema file")

	schema, err := request.ParseRequestSchema(data)
	utils.CheckError(err, "parsing request schema")

	req, err := request.CreateRequest(schema)
	utils.CheckError(err, "creating request")

	start := time.Now()
	resp, err := request.ExecuteRequest(req)
	utils.CheckError(err, "executing request")
	elapsed := time.Since(start)

	printResponse(args, resp, elapsed)
}

func printResponse(args arguments.CommandLineArguments, resp *http.Response, elapsed time.Duration) {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	utils.CheckError(err, "reading response body")

	statusCode := http_utils.Status(resp.StatusCode)
	statusColor := color.GetHTTPStatusColors(statusCode.Code())
	statusStr := fmt.Sprintf("%d (%s)", resp.StatusCode, statusCode.String())

	elapsedColor := color.GetExecutionTimeColor(elapsed.Milliseconds())
	elapsedStr := elapsed.String()

	var output []byte

	if strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		output, err = utils.PrettyJSON(body)

		if err != nil {
			output = body
		}
	} else {
		output = body
	}

	if args.OutputFile != "" {
		err := os.WriteFile(args.OutputFile, output, 0644)
		utils.CheckError(err, "writing response to file")
		return
	}

	fmt.Println(string(utils.LimitLines(output, args.NumberOfLines)))

	fmt.Printf("Request finished with status code %s in %s\n",
		color.Colorize(statusStr, statusColor),
		color.Colorize(elapsedStr, elapsedColor))
}
