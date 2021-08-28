package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// TODO don't search files that aren't text files

func Input_Depth(depth *int, err error, do bool) {
	// if do is true then get user input
	// while err or depth < 0 get user input
	var depth_s string
	if do {
		fmt.Print("Depth: ")
		fmt.Scanln(&depth_s)
		*depth, err = strconv.Atoi(depth_s)
	}
	for err != nil || *depth < 0 {
		fmt.Println("Depth must be an int >= 0")
		fmt.Print("Depth: ")
		fmt.Scanln(&depth_s)
		*depth, err = strconv.Atoi(depth_s)
	}
}

func main() {
	start := time.Now()
	var dir, pattern string
	var depth int
	var err error
	fmt.Print("\nDirectory: ")
	if len(os.Args) > 1 {
		dir = os.Args[1]
		dir, err = get_dir(dir)
		fmt.Println(dir)
	} else {
		fmt.Scanln(&dir)
	}
	dir, err = get_dir(dir)
	if err != nil {
		fmt.Println("Error in getting directory")
		return
	}

	fmt.Print("Pattern: ")
	if len(os.Args) > 2 {
		pattern = os.Args[2]
		fmt.Println(pattern)
	} else {
		fmt.Scanln(&pattern)
	}

	if len(os.Args) > 3 {
		depth, err = strconv.Atoi(os.Args[3])
		fmt.Println("Depth:", depth)
		// while error or less than 0 then gets input
		Input_Depth(&depth, err, false)
	} else {
		Input_Depth(&depth, err, true)
	}
	fmt.Println(time.Since(start))
	fmt.Print("\n\n")
	search_dir(dir, pattern, depth, true, start)
	// fmt.Println(dir, pattern, depth)
}
