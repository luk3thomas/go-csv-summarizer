package main

import (
	"flag"
	"fmt"
)

func main() {
	var col string
	var file string

	flag.StringVar(&file, "file", "", "Path to a CSV file")
	flag.StringVar(&col, "column", "1", "Numerical columns")

	flag.Parse()

	options := BuildOptions(
		file,
		col,
	)

	summary := Summarize(options)

	fmt.Printf("%+v\n", summary)
}
