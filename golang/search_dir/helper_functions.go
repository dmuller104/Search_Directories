package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// func search_dir_surface(dir string, pattern string) {
// 	search_dir(dir, pattern, 0, true)
// }

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Line struct {
	number int
	line   []byte
}

type FileResult struct {
	name  string
	lines []Line
}

func check_file_type(name string) int {
	switch filepath.Ext(name) {
	case ".txt":
		return 1
	case ".csv":
		return 2
	case ".cpp":
		return 3
	case ".py":
		return 3
	case ".go":
		return 3
	case ".html":
		return 3
	case ".js":
		return 3
	case ".clj":
		return 3
	default:
		return -1
	}
}

func get_dir(input string) (string, error) {
	if len(strings.Trim(input, " ")) == 0 {
		input = "."
	}
	return filepath.Abs(input)
}

func search_file(name string, pattern string, pipe chan *FileResult) {
	if check_file_type(name) == -1 {
		// fmt.Println("-1:", name)
		pipe <- nil
		return
	}
	var file FileResult
	file.name = name
	dat, err := ioutil.ReadFile(file.name)
	check(err)
	lines := bytes.Split(dat, []byte{'\n'})
	for i, line := range lines {
		r, _ := regexp.Compile(pattern)
		if r.Find(line) != nil {
			var l Line
			l.line = line
			l.number = i
			file.lines = append(file.lines, l)
		}
	}
	if len(file.lines) == 0 {
		pipe <- nil
	} else {
		pipe <- &file
	}
}

func search_dir(dir string, pattern string, depth int, show bool, start time.Time) {
	fmt.Println(time.Since(start))
	// fmt.Println("\nSearch Directory")
	// fmt.Println(dir)
	files, err := ioutil.ReadDir(dir)
	var dirs []string
	check(err)
	pipe := make(chan *FileResult)
	var num_searched = 0
	for _, file := range files {
		if file.IsDir() {
			// fmt.Println("[DIR]", path.Join(dir, file.Name()))
			dirs = append(dirs, file.Name())
		} else {
			go search_file(path.Join(dir, file.Name()), pattern, pipe)
			num_searched++
		}
	}
	for i := 0; i < num_searched; i++ {
		results := <-pipe
		if results != nil {
			fmt.Printf("%s\n", results.name)
			for _, line := range results.lines {
				if show {
					fmt.Printf("    %d %+s\n", line.number, string(line.line))
				}
			}
		}
	}
	if depth > 0 {
		for _, d := range dirs {
			search_dir(path.Join(dir, d), pattern, depth-1, show, start)
			// fmt.Println("Search:", d)
		}
	}

}
