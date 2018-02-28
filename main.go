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

	// use recursion to scan through each token
	InterpreteProgram(bytes)
}

// create empty int variable array
var memory = make([]int, 512)
var ptr = 0

// InterpreteProgram runs said cases recursively if memory[ptr] != 0
func InterpreteProgram(bytes []byte) {
	loop := []byte{}
	looping := false
	jump2end := false

	for _, ch := range bytes {
		s := string(ch)

		if s != "]" && jump2end {
			continue
		}

		if looping {
			loop = append(loop, ch)
		}

		switch s {
		case ">":
			ptr++
		case "<":
			if ptr != 0 {
				ptr--
			}
		case ".":
			fmt.Printf("%s", string(memory[ptr]))
		case "+":
			if memory[ptr] == 255 {
				memory[ptr]--
			}
			memory[ptr]++
		case "-":
			if memory[ptr] == 0 {
				memory[ptr] = 256
			}
			memory[ptr]--
		case "[":
			if memory[ptr] == 0 {
				jump2end = true
				continue
			}

			looping = true
			loop = append([]byte{}, ch)
		case "]":
			if jump2end {
				jump2end = false
			}

			if memory[ptr] != 0 {
				InterpreteProgram(loop)
			}

			looping = false
		}
	}
}
