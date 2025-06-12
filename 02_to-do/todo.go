package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
)

// Have to be exported in order to be json.Marshalled
type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type list []item

func (l *list) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := " "
		if t.Done {
			prefix = "X "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}

	return formatted
}

func (l *list) add(task string) {
	t := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *list) complete(i int) error {
	ls := *l
	if i < 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	ls[i-1].Done = true
	ls[i-1].CompletedAt = time.Now()

	return nil
}

func (l *list) delete(i int) error {
	ls := *l
	if i < 0 || i > len(ls) {
		return fmt.Errorf("item %d does not exist", i)
	}

	// *l = append(ls[:i-1], ls[i:]...)
	*l = slices.Delete(*l, i-1, i)
	return nil
}

func (l *list) save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	fmt.Println((*l))
	fmt.Print(js)

	return os.WriteFile(filename, js, 0644)
}

func (l *list) load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)

}
