package exercises

// task: get the url and filename from the CLI args, fetch the html file and store in a html file
/*
	- error handling: fetching, decoding, file creation and writing to the file
	- timeout for fetch calls
	- generting a safe filename and writing to that file
*/

// get the flag url and filename rather than looping for all th CLI
import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"
)

var fileExtension string = ".html"
var outDir string = "./outDir/"

func FetchFromURL() {

	urlFlag := flag.String("url", "", "URL to fetch")

	flag.Parse()
	if *urlFlag == "" {
		fmt.Fprintf(os.Stderr, "cli arg: %v\n", "Please input a --url flag")
		os.Exit(1)
	}

	if !isValidURL(*urlFlag) {
		fmt.Fprintf(os.Stderr, "client error: %v\n", "Invalid URL")
		os.Exit(1)
	}

	res, err := http.Get(*urlFlag)

	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if err := identifyAndLogError(res.StatusCode); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fileName := getFileName(*urlFlag)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file  creation: %v\n", err)

	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "write the response: %v\n", err)
		os.Exit(1)

	}

}

func identifyAndLogError(statusCode int) error {
	switch statusCode {
	case http.StatusOK:
		return nil // No error for OK status
	case http.StatusNotFound:
		return fmt.Errorf("error 404: resource not found")
	case http.StatusInternalServerError:
		return fmt.Errorf("error 500: internal server error")
	default:
		return fmt.Errorf("unexpected status code %d", statusCode)
	}
}
func getFileName(urlString string) string {
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return "error.html"
	}

	host := parsedURL.Hostname()
	_path := parsedURL.Path
	dirPath := path.Join(outDir, host)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		fmt.Fprintf(os.Stderr, "error creating directory: %v\n", err)
		return "error.html"
	}
	var filePath string
	if sanitizeFileName(_path) == "" {
		filePath = path.Join(dirPath, fmt.Sprintf("%s%s", host, fileExtension))

	} else {

		filePath = path.Join(dirPath, fmt.Sprintf("%s_%s%s", host, sanitizeFileName(_path), fileExtension))
	}

	return filePath

}

func sanitizeFileName(name string) string {
	invalidChars := regexp.MustCompile("[[:^alnum:]]")
	name = invalidChars.ReplaceAllString(name, "_")
	name = strings.Trim(name, "_")
	return name
}

func isValidURL(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)
	return err == nil
}
