package main

import (
	"os"
	"testing"
)

func TestAdd(t *testing.T) {
	l := list{}

	taskName := "New Task"
	l.add(taskName)

	if l[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l[0].Task)
	}

	if l[0].Done {
		t.Errorf("New Task should not be completed.")
	}

	l.complete(1)
	if !l[0].Done {
		t.Errorf("New Task should be completed.")
	}
}

func TestDelete(t *testing.T) {
	l := list{}

	tasks := []string{
		"New Task 1",
		"New Task 2",
		"New Task 3",
	}

	for _, v := range tasks {
		l.add(v)
	}

	if l[0].Task != tasks[0] {
		t.Errorf("Expected %q, got %q instead.", tasks[0], l[0].Task)
	}

	l.delete(2)

	if len(l) != 2 {
		t.Errorf("Exected list length %d, got %d", 2, len(l))
	}

	if l[1].Task != tasks[2] {
		t.Errorf("Expected %q, got %q instead.", tasks[2], l[1].Task)
	}

}

func TestSaveLoad(t *testing.T) {
	l1 := list{}
	l2 := list{}

	taskName := "New Task"

	l1.add(taskName)

	if l1[0].Task != taskName {
		t.Errorf("Expected %q, got %q instead.", taskName, l1[0].Task)
	}

	tf, err := os.CreateTemp("", "")

	if err != nil {
		t.Fatalf("Error creating temp file: %s", err)
	}

	defer os.Remove(tf.Name())

	if err := l1.save(tf.Name()); err != nil {
		t.Fatalf("Error saving list to file: %s", err)
	}

	if err := l2.load(tf.Name()); err != nil {
		t.Fatalf("Error loading list from file: %s", err)
	}

	if l1[0].Task != l2[0].Task {
		t.Errorf("Task %q should match %q Task.", l1[0].Task, l2[0].Task)
	}
}
