package main

import (
	"fmt"
)

func main() {
	var input_array [256]bool
	var key [256]bool

	var output_array [256]bool
	for i:=0; i<len(input_array); i++ {
		output_array[i] = input_array[i] == key[i]
	}
	fmt.Println(output_array[0]);
}
