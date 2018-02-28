package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// get filename from cmd args
	fn := os.Args[1]

	// open file
	fd, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	// put file into string
	reader := bufio.NewReader(fd)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	// prints hello.bf as bytes
	fmt.Println(bytes)
}
