package main

import (
	"flag"
	"fmt"
	"os"
)

const todoFileName = ".todo.json"

func main() {

	task := flag.String("task", "", "Task to be included in the ToDo list")
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
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}

	case *complete > 0:

		if err := l.complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case *task != "":
		l.add(*task)

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
