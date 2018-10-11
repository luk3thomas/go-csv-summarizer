package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Column type
type Column int64

// DateFormat format for a date
type DateFormat int

const (
	// DateFormatYear format by year
	DateFormatYear DateFormat = 0
	// DateFormatMonth format by month
	DateFormatMonth DateFormat = 1
)

// Options struct
type Options struct {
	file    *csv.Reader
	column  Column
	datecol Column
	datefmt string
	dateagg DateFormat
}

// BuildOptions create options for parser
func BuildOptions(file string, col string, dcol string, dfmt string, dagg string) Options {
	datecol := getColumnOption(dcol)

	if datecol == 0 {
		if dfmt != "" || dagg != "" {
			fmt.Println("date-column is required")
			os.Exit(1)
		}
	}

	if dfmt == "" && datecol != 0 {
		fmt.Println("date-format is required")
		os.Exit(1)
	}

	if dcol == "" && datecol != 0 {
		fmt.Println("date-column is required")
		os.Exit(1)
	}

	opt := Options{
		file:    getFileReader(file),
		column:  getColumnOption(col),
		datecol: datecol,
		datefmt: dfmt,
		dateagg: getDateAggregation(dagg),
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

func getColumnOption(col string) Column {
	i, err := strconv.ParseInt(col, 10, 64)
	if err != nil {
		return Column(0)
	}
	return Column(i)
}

func getDateAggregation(dfmt string) DateFormat {
	switch strings.ToLower(dfmt) {
	case "year":
		return DateFormatYear
	default:
		return DateFormatMonth
	}
}
