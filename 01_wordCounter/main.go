package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func count(r io.Reader, lines bool, bytes bool) string {
	scanner := bufio.NewScanner(r)

	scanner.Split(bufio.ScanLines)
	wc := 0
	lc := 0
	bc := 0
	for scanner.Scan() {
		line := scanner.Text()
		lc++
		wc += len(strings.Fields(line))
		bc += len(scanner.Bytes())
	}

	out := fmt.Sprintf("Words: %d\n", wc)

	if lines {
		out += fmt.Sprintf("Lines: %d\n", lc)
	}
	if bytes {
		out += fmt.Sprintf("Bytes: %d\n", bc)
	}

	return out
}

func main() {

	lines := flag.Bool("l", false, "Count lines")
	bytes := flag.Bool("b", false, "Count bytes")

	flag.Parse()

	fmt.Println(count(os.Stdin, *lines, *bytes))
}
