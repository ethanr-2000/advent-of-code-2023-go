package main

import (
	"advent-of-code-go/pkg/list"
	_ "embed"
	"flag"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
}

func main() {
	var part int
	flag.IntVar(&part, "part", 0, "part 1 or 2")
	flag.Parse()

	if part == 1 {
		ans := part1(input)
		fmt.Println("Running part 1")
		clipboard.WriteAll(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else if part == 2 {
		ans := part2(input)
		fmt.Println("Running part 2")
		clipboard.WriteAll(fmt.Sprintf("%v", ans))
		fmt.Println("Output:", ans)
	} else {
		fmt.Println("Running all")
		ans1 := part1(input)
		fmt.Println("Part 1 Output:", ans1)
		ans2 := part2(input)
		fmt.Println("Part 2 Output:", ans2)
	}
}

func part1(input string) int {
	files := getFiles(input)

	return checksum(compactFiles(files))
}

func part2(input string) int {
	files := getFiles(input)

	return checksum(compactFilesWhole(files))
}

type File struct {
	id                    int
	length                int
	subsequentSpaceLength int
	moved                 bool
}


// Helper functions for part 1

func getFiles(filesystem string) []File {
	var files []File
	for i := 0; i < len(filesystem); i += 2 {
		val, _ := strconv.Atoi(string(filesystem[i]))

		spaceLength := 0
		if i != len(filesystem)-1 {
			// if we're not at the end, there is space after
			spaceLength, _ = strconv.Atoi(string(filesystem[i+1]))
		}

		files = append(files, File{
			id:                    i / 2,
			length:                val,
			subsequentSpaceLength: spaceLength,
		})
	}
	return files
}

func compactFiles(files []File) []File {
	for filesNotOptimallyCompact(files) {
		files = moveLastFileToEmptySpace(files)
	}

	return files
}

func moveLastFileToEmptySpace(files []File) []File {
	for i, f := range files[0 : len(files)-1] {
		if f.subsequentSpaceLength > 0 {
			if f.subsequentSpaceLength >= files[len(files)-1].length {
				// there is enough or more than enough space

				// the penultimate file is the last file, and absorbs the last file's length and space
				files[len(files)-2].subsequentSpaceLength += files[len(files)-1].length + files[len(files)-1].subsequentSpaceLength

				files = slices.Insert[[]File](files, i+1, File{
					id:                    files[len(files)-1].id,
					length:                files[len(files)-1].length,
					subsequentSpaceLength: f.subsequentSpaceLength - files[len(files)-1].length,
				})

				files[i].subsequentSpaceLength = 0 // the potential spaces are now part of the next file

				return files[0 : len(files)-1] // remove the last file, we just moved it to an empty space
			} else if f.subsequentSpaceLength < files[len(files)-1].length {
				// there is not enough space - fill the space but keep process the last file

				files[len(files)-1].subsequentSpaceLength += f.subsequentSpaceLength
				files[len(files)-1].length -= f.subsequentSpaceLength

				files = slices.Insert[[]File](files, i+1, File{
					id:                    files[len(files)-1].id,
					length:                f.subsequentSpaceLength,
					subsequentSpaceLength: 0,
				})

				files[i].subsequentSpaceLength = 0 // the potential spaces are now part of the next file
				// carry on spreading the file. it still exists
			}
		}
	}
	return files
}

func filesNotOptimallyCompact(files []File) bool {
	// the last file is the only one allowed to have spaces
	for _, f := range files[0 : len(files)-1] {
		if f.subsequentSpaceLength > 0 {
			return true
		}
	}
	return false
}

func checksum(files []File) int {
	positionId := -1
	c := 0
	for _, f := range files {
		for i := 0; i < f.length; i++ {
			positionId++
			c += positionId * f.id
		}
		positionId += f.subsequentSpaceLength
	}
	return c
}

// Helper functions for part 2

func compactFilesWhole(files []File) []File {
	for i := len(files) - 1; i >= 0; i-- {
		if !files[i].moved {
			newFiles, moved := attemptMoveFileToSpace(files, i)
			files = newFiles

			if moved {
				// without this, we've skipped a file!
				i ++
			}
		}
	}
	return files
}

func attemptMoveFileToSpace(files []File, indexOfFileToMove int) ([]File, bool) {
	for i := range files[0:indexOfFileToMove] {
		if files[i].subsequentSpaceLength >= files[indexOfFileToMove].length {
			// there is enough or more than enough space
			// the file before the one moving absorbs the last file's length and space
			files[indexOfFileToMove-1].subsequentSpaceLength += files[indexOfFileToMove].length + files[indexOfFileToMove].subsequentSpaceLength

			files = slices.Insert[[]File](files, i+1, File{
				id:                    files[indexOfFileToMove].id,
				length:                files[indexOfFileToMove].length,
				subsequentSpaceLength: files[i].subsequentSpaceLength - files[indexOfFileToMove].length,
			})

			files[i].subsequentSpaceLength = 0 // the potential spaces are now part of the next file

			return list.DeleteAtIndices[File](files, []int{indexOfFileToMove + 1}), true // we've added a file, so the index is one off
		}
	}
	return files, false
}
