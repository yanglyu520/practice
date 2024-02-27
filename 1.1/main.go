// Exercis e 1.1: Modify the echo program to also print os.Args[0], the name of the command
// that invoked it.
// Exercis e 1.2: Modify the echo program to print the index and value of each of its arguments,
// one per line.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	echo1(os.Args)
	echo2(os.Args)
}

func echo1(stringSlice []string) {
	var s, sep string
	for _, comp := range stringSlice[1:] {
		s += sep + comp
		sep = " "
	}
	fmt.Println(s)
}

func echo2(stringSlice []string) {
	var s, sep string
	for _, comp := range stringSlice {
		s += sep + comp
		sep = "\n"
	}

	fmt.Println(s)
}

func echo3(stringSlice []string) {
	strings.Join(stringSlice, "\n")
}
