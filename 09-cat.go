package main

import (
	"flag"
	"fmt"
	"os"
)

var numberFlag = flag.Bool("n", false, "number each line")

func main() {
	flag.Parse()
	if flag.NArg() == 1 {

	}
	fmt.Print(os.Stdout)
}