// Package skeleton makes skeletons to be filled out with solutions.
package skeleton

import (
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"text/template"
)

//go:embed template/*.go
var fs embed.FS

// Run makes a skeleton main.go and main_test.go file for the given day and year
func Run(day, year int) {
	if day > 25 || day <= 0 {
		log.Fatalf("invalid -day value, must be 1 through 25, got %v", day)
	}

	if year < 2015 {
		log.Fatalf("year is before 2015: %d", year)
	}

	ts, err := template.ParseFS(fs, "template/*.go")
	if err != nil {
		log.Fatalf("parsing template directory: %s", err)
	}

	mainFilename := filepath.Join(dirname(), "../../", fmt.Sprintf("%d/day%02d/main.go", year, day))
	testFilename := filepath.Join(dirname(), "../../", fmt.Sprintf("%d/day%02d/main_test.go", year, day))
	inputFilename := filepath.Join(dirname(), "../../", fmt.Sprintf("%d/day%02d/input.txt", year, day))

	err = os.MkdirAll(filepath.Dir(mainFilename), os.ModePerm)
	if err != nil {
		log.Fatalf("making directory: %s", err)
	}

	ensureNotOverwriting(mainFilename)
	ensureNotOverwriting(testFilename)
	ensureNotOverwriting(inputFilename)

	mainFile, err := os.Create(mainFilename)
	if err != nil {
		log.Fatalf("creating main.go file: %v", err)
	}
	testFile, err := os.Create(testFilename)
	if err != nil {
		log.Fatalf("creating main_test.go file: %v", err)
	}
	_, err = os.Create(inputFilename)
	if err != nil {
		log.Fatalf("creating input.txt file: %v", err)
	}

	ts.ExecuteTemplate(mainFile, "main.go", nil)
	ts.ExecuteTemplate(testFile, "main_test.go", nil)
	fmt.Printf("templates made for %d-day%d\n", year, day)
}

// Dirname is a port of __dirname in node
func dirname() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("getting calling function")
	}
	return filepath.Dir(filename)
}

func ensureNotOverwriting(filename string) {
	_, err := os.Stat(filename)
	if err == nil {
		log.Fatalf("File already exists: %s", filename)
	}
}
