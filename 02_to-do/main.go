package main

import (
	"fmt"
	"os"
	"strings"
)

const todoFileName = ".todo.json"

func main() {
	l := &list{}

	if err := l.load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {

	case len(os.Args) == 1:
		for _, item := range *l {
			fmt.Println(item.Task)
		}

	default:
		item := strings.Join(os.Args[1:], "")
		l.add(item)

		if err := l.save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}
