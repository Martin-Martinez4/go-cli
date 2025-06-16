package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

type config struct {
	// extensions to filter out
	ext string

	// min file size
	size int64

	// list files
	list bool

	del bool

	wLog io.Writer

	archiveDir string
}

func run(root string, out io.Writer, cfg config) error {
	delogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)

	return filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filterOut(path, cfg.ext, cfg.size, info) {
				return nil
			}

			if cfg.list {
				return listFile(path, out)
			}
			if cfg.archiveDir != "" {
				if err := archiveFile(cfg.archiveDir, root, path); err != nil {
					return err
				}
			}

			if cfg.del {
				return deleteFile(path, delogger)
			}

			return listFile(path, out)
		})
}

func main() {
	root := flag.String("root", ".", "Root directory to start")
	logFile := flag.String("log", "", "Log deletes to this file")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	archiveDir := flag.String("archiveDir", "", "Archive directory")
	del := flag.Bool("del", false, "Delete files")
	flag.Parse()

	var (
		f   = os.Stdout
		err error
	)

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	c := config{
		ext:        *ext,
		size:       *size,
		list:       *list,
		del:        *del,
		wLog:       f,
		archiveDir: *archiveDir,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
