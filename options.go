package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Column type
type Column int64

// Options struct
type Options struct {
	file   *csv.Reader
	column Column
}

// BuildOptions create options for parser
func BuildOptions(file string, col string) Options {
	opt := Options{
		file:   getFileReader(file),
		column: getColumn(col),
	}
	return opt
}

func getFileReader(file string) *csv.Reader {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	// defer f.Close()
	return csv.NewReader(f)
}

func getColumn(col string) Column {
	i, err := strconv.ParseInt(col, 10, 64)
	if err != nil {
		fmt.Printf("Parse error %v", err)
		os.Exit(1)
	}
	return Column(i)
}
