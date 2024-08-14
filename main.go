package main

import (
	"flag"
	"github.com/mburbidg/grmtest/parser"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filename := flag.String("file", "", "Filename to read GQL-program from")
	dirname := flag.String("dir", "", "Directory name containing GQL-programs")
	flag.Parse()
	if *filename != "" {
		parseFile(*filename)
	}
	if *dirname != "" {
		parseDir(*dirname)
	}
}

func parseFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed to open '%s': %v", filename, err)
	}
	b, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read '%s': %v", filename, err)
	}
	err = parser.Parse(string(b))
	if err != nil {
		log.Printf("Syntax error in '%s': %v", filename, err)
	}
	log.Printf("Success: %s", filename)
}

func parseDir(dirname string) {
	files, err := os.ReadDir(dirname)
	if err != nil {
		log.Fatalf("Failed to read directory '%s': %v", dirname, err)
	}

	for _, file := range files {
		path := filepath.Join(dirname, file.Name())
		if !file.IsDir() {
			if filepath.Ext(path) == ".gql" {
				parseFile(path)
			}
		} else {
			parseDir(path)
		}
	}
}
