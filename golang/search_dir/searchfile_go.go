package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func search_file(name string, pattern string) ([]int, [][]byte) {
	var line_matches [][]byte
	var line_numbers []int
	dat, err := ioutil.ReadFile(name)
	check(err)
	lines := bytes.Split(dat, []byte{'\n'})
	for i, line := range lines {
		// fmt.Println(line)
		r, _ := regexp.Compile(pattern)
		if r.Find(line) != nil {
			line_matches = append(line_matches, line)
			line_numbers = append(line_numbers, i)
			// fmt.Println(i, string(line))
		}
	}

	return line_numbers, line_matches
	// fmt.Println(dat)
}

func do_search_file() {
	fmt.Println("\nFirst method, not goroutinable")
	line_matches, matches := search_file("testDir\\searchable.txt", "[hH]ello")
	for i, match := range matches {
		fmt.Println(line_matches[i], string(match))
	}
	line_matches, matches = search_file("testDir\\searchable2.txt", "[hH]ello")
	for i, match := range matches {
		fmt.Println(line_matches[i], string(match))
	}
}

type Line struct {
	number int
	line   []byte
}

func search_file_go(name string, pattern string, match_chan chan []Line) {
	var line_matches []Line
	// var line_numbers []int
	dat, err := ioutil.ReadFile(name)
	check(err)
	lines := bytes.Split(dat, []byte{'\n'})
	for i, line := range lines {
		// fmt.Println(line)
		r, _ := regexp.Compile(pattern)
		if r.Find(line) != nil {
			var l Line
			l.line = line
			l.number = i
			line_matches = append(line_matches, l)
			// line_matches = append(line_matches, line)
			// line_numbers = append(line_numbers, i)
			// fmt.Println(i, string(line))
		}
	}

	// line_chan <- line_numbers
	match_chan <- line_matches
	// return line_numbers, line_matches
	// fmt.Println(dat)
}

func do_search_file_go() {
	fmt.Println("\n----Go routines----")
	match_chan := make(chan []Line)
	// line_chan := make(chan []int)
	go search_file_go("testDir\\searchable.txt", "[hH]ello", match_chan)
	go search_file_go("testDir\\searchable2.txt", "[hH]ello", match_chan)

	// mat := <-match_chan
	var l_matches []Line
	l_matches = <-match_chan
	for _, line := range l_matches {
		fmt.Println(line.number, string(line.line))
	}
	l_matches = <-match_chan
	for _, line := range l_matches {
		fmt.Println(line.number, string(line.line))
	}
}

type FileResult struct {
	name  string
	lines []Line
}

func search_file_go2(name string, pattern string, pipe chan *FileResult) {
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

func display_with_pipe(result *FileResult) {
	fmt.Println(result.name)
	fmt.Printf("%+v\n", result.lines)
}

func do_search_file_go2() {
	fmt.Println("\n----Go Routines but smarter----")
	pipe := make(chan *FileResult)
	go search_file_go2("testDir\\searchable.txt", "[hH]ello", pipe)
	go search_file_go2("testDir\\searchable2.txt", "[hH]ello", pipe)
	var result *FileResult
	result = <-pipe
	display_with_pipe(result)
	result = <-pipe
	display_with_pipe(result)
	// fmt.Printf("%+v\n", *<-pipe)
	// fmt.Printf("%+v\n", *<-pipe)
}

func do_search_list_of_files() {

	fmt.Println("\n----Go Routines but smarter with list----")
	var files []string
	files = append(files, "testDir\\searchable.txt")
	files = append(files, "testDir\\searchable2.txt")
	pipe := make(chan *FileResult)
	pattern := "[hH]ello"
	for _, file := range files {
		go search_file_go2(file, pattern, pipe)
	}
	var results []*FileResult
	for i := 0; i < len(files); i++ {
		result := <-pipe
		results = append(results, result)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].name < results[j].name
	})
	for _, element := range results {
		fmt.Printf("%s\n%+v\n", element.name, element.lines)
	}
}

func get_dir(input string) (string, error) {
	if len(strings.Trim(input, " ")) == 0 {
		input = "."
	}
	return filepath.Abs(input)
}

func working_with_path(rel_path string) string {
	// take relative path as parameter
	// working_dir, _ := os.Getwd()
	working_dir, _ := get_dir("testDir")
	dir := path.Join(working_dir, rel_path)
	return dir
	// return full path

}

func search_list_of_paths() {
	fmt.Println("\n----Search all files in current directory----")
}

func search_dir_f(dir string) {
	fmt.Println("\n----Search Directory----")
	files, err := ioutil.ReadDir(dir)
	check(err)
	for _, file := range files {
		if file.IsDir() {
			fmt.Print("[DIR] ")
		} else {
			fmt.Print("      ")
		}
		fmt.Println(path.Join(dir, file.Name()))
	}
}

func do_search_dir_f() {
	fmt.Println("\nFirst search dir")
	dir, err := get_dir("..\\search_dir\\testDir")
	check(err)
	search_dir_f(dir)
}

func search_dir_surface(dir string, pattern string) {
	search_dir(dir, pattern, 0)
}
func search_dir(dir string, pattern string, depth int) {
	fmt.Println("\nSearch Directory")
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
			go search_file_go2(path.Join(dir, file.Name()), pattern, pipe)
			num_searched++
		}
	}
	for i := 0; i < num_searched; i++ {
		results := <-pipe
		if results != nil {
			fmt.Printf("%s\n", results.name)
			for _, line := range results.lines {
				fmt.Printf("%d %+s\n", line.number, string(line.line))
			}
		}
	}
	if depth > 0 {
		for _, d := range dirs {
			search_dir(path.Join(dir, d), pattern, depth-1)
			// fmt.Println("Search:", d)
		}
	}

}

func do_search_dir() {
	fmt.Println("\n----Search files in directory----")
	dir, err := get_dir("..\\search_dir\\testDir")
	check(err)
	search_dir_surface(dir, "[hH]ello")
}
func do_search_dirs() {
	fmt.Println("\n----Search files in directories----")
	dir, err := get_dir("..\\search_dir\\testDir")
	check(err)
	search_dir(dir, "[hH]ello", 1)
}

func main() {
	fmt.Println("----Start----")
	do_search_file()
	do_search_file_go()
	do_search_file_go2()
	do_search_list_of_files()
	dir_unof := working_with_path("searchable.txt")
	res, _ := search_file(dir_unof, "[hH]ello")
	fmt.Printf("%+v\n", res)
	do_search_dir_f()
	do_search_dir()
	do_search_dirs()
	fmt.Println("----Complete----")

}
