package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		arg := []rune(args[i])
		for j := 0; j < len(arg); j++ {
			z01.PrintRune(arg[j])
		}
		z01.PrintRune('\n')
	}
}
