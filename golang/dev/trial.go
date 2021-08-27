package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"
)

// "path/filepath"

// Trial 1 - structs and goroutines
type person struct {
	name string
	age  int
}

func change_person(p *person) {
	p.age = 5
	p.name = "Collin"
}

func trial1() {
	fmt.Println("\nTrial 1")
	fmt.Println("Are structs changed in goroutines?")
	var p person
	p.name = "Derek"
	p.age = 23
	fmt.Printf("%+v\n", p)
	go change_person(&p)
	fmt.Printf("%+v\n", p)
	time.Sleep(1)
	fmt.Printf("%+v\n", p)
}

// Trial 2 - create struct in goroutine
func create_person(pipe chan *person) {
	var p person
	p.age = 9
	p.name = "Calvin"
	pipe <- &p
}

func trial2() {
	fmt.Println("\nTrial 2")
	fmt.Println("Return variable by pointer in goroutine")
	pipe := make(chan *person)
	go create_person(pipe)
	var p person
	p = *<-pipe
	fmt.Printf("%+v\n", p)
}

func trial3() {
	fmt.Println("\nTrial 3")
	fmt.Println("Using full file paths and libraries")
	readdir, _ := os.ReadDir("text.txt")
	fmt.Println("\nReadDir")
	fmt.Println(readdir)
	fmt.Printf("%T\n", readdir)
	// absfilepath, _ := filepath.Abs("text.txt")
	absfilepath, _ := filepath.Abs(".")
	fmt.Println("\nabsfilepath")
	fmt.Println(absfilepath)
	fmt.Printf("%T\n", absfilepath)
	workingdir, _ := os.Getwd() // gives dir of location the program is called from
	fmt.Println("\nworkingdir")
	fmt.Println(workingdir)
	fmt.Printf("%T\n", workingdir)
	programdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Println("\nprogramdir")
	fmt.Println(programdir)
	fmt.Printf("%T\n", programdir)
	executable, _ := os.Executable()
	fmt.Println("\nexecutable")
	fmt.Println(executable)
	fmt.Printf("%T\n", executable)
	join_wd := path.Join(workingdir, "file.txt") // gives dir of location the program is called from
	fmt.Println("\njoin_wd")
	fmt.Println(join_wd)
	fmt.Printf("%T\n", join_wd)
}

/////////////

func trial4() {
	fmt.Println("\nTrial 4")
	fmt.Println("Get list of files inside of directory")
	dir, _ := filepath.Abs(".")
	files, _ := ioutil.ReadDir(dir)
	for _, file := range files {
		if file.IsDir() {
			fmt.Print("\t")
		}
		fmt.Println(file.Name())
		fmt.Println(path.Join(dir, file.Name()))
	}
}

func main() {
	fmt.Println("\n----TRIAL----")
	// trial1()
	// trial2()
	trial3()
	trial4()
	fmt.Println("\n----COMPLETE----")
}
