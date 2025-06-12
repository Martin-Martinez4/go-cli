package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

var todoFileName = ".todo.json"

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	s := bufio.NewScanner(r)
	s.Scan()
	if err := s.Err(); err != nil {
		return "", err
	}

	if len(s.Text()) == 0 {
		return "", fmt.Errorf("Task cannot be blank")
	}

	return s.Text(), nil
}

func main() {

	if os.Getenv("TODO_FILENAME") != "" {

		todoFileName = os.Getenv("TODO_FILENAME")
	}

	add := flag.Bool("add", false, "Task to be included in the ToDo list")
	lst := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	flag.Parse()

	l := &list{}

	if err := l.load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)

	}

	switch {

	case *lst:
		fmt.Print(l)

	case *complete > 0:

		if err := l.complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *add:
		t, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		l.add(t)

		if err := l.save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		// invalid flag
		fmt.Fprintln(os.Stderr, "Invalid flag")
		os.Exit(1)

	}
}
