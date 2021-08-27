package main

import (
	"fmt"
	"os"
	"strconv"
)

func Input_Depth(depth *int, err error, do bool) {
	var depth_s string
	// fmt.Print("Depth: ", *depth)
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
		// fmt.Println("error: ", err)
		Input_Depth(&depth, err, false)
		// for err != nil || depth < 0{
		// 	fmt.Println("Depth must be an int >= 0")
		// 	fmt.Print("Depth: ")
		// 	fmt.Scanln(&depth_s)
		// 	depth,err = strconv.Atoi(depth_s)
		// }
	} else {
		Input_Depth(&depth, err, true)
		// fmt.Print("Depth: ")
		// fmt.Scanln(&depth_s)
		// depth,err := strconv.Atoi(depth_s)
		// for err != nil || depth < 0{
		// 	fmt.Println("Depth must be an int >= 0")
		// 	fmt.Print("Depth: ")
		// 	fmt.Scanln(&depth_s)
		// 	depth,err = strconv.Atoi(depth_s)
		// }
	}

	fmt.Print("\n\n")
	search_dir(dir, pattern, depth)
	// fmt.Println(dir, pattern, depth)
}
