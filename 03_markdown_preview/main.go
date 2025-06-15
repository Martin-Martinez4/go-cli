package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const header = `<!DOCTYPE html>
<html>
	<head>
		<meta http-equiv="content-type" content="text/html" charset="utf-8" >
		<title>Markdown Preview Tool</title>
	</head>
	<body>
`

const footer = `
	</body>
</html>`

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var buffer bytes.Buffer
	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHTML(outfilename string, data []byte) error {
	return os.WriteFile(outfilename, data, 0644)
}

func run(path string) error {
	input, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	htmlData := parseContent(input)
	outName := fmt.Sprintf("%s.html", filepath.Base(path))
	fmt.Println(outName)
	return saveHTML(outName, htmlData)
}

func main() {
	path := flag.String("path", "", "Markdown file to preview")
	flag.Parse()

	if *path == "" {
		// flag usage error
		flag.Usage()
		os.Exit(1)
	}
	if err := run(*path); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
