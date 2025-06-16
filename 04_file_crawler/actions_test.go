package main

import (
	"os"
	"testing"
)

func TestFilterOut(t *testing.T) {

	testCases := []struct {
		name     string
		file     string
		ext      string
		minSize  int64
		expected bool
	}{
		{"Filter No Extenstion", "testdata/dir.log", "", 0, false},
		{"Filter Extension Match", "testdata/dir.log", ".log", 0, false},
		{"Filter Extension No Match", "testdata/dir.log", ".sh", 0, true},
		{"Filter Extension Size Match", "testdata/dir.log", ".log", 10, false},
		{"Filter Extension Size No Match", "testdata/dir.log", ".log", 20, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			info, err := os.Stat(tc.file)
			if err != nil {
				t.Fatal(err)
			}

			f := filterOut(tc.file, tc.ext, tc.minSize, info)

			if f != tc.expected {
				t.Errorf("Expected '%t', got '%t' instead\n", tc.expected, f)
			}
		})
	}
}
