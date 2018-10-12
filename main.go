package main

import (
	"flag"
	"fmt"
)

func main() {
	var col string
	var file string
	var datecol string
	var datefmt string
	var dateagg string

	flag.StringVar(&file, "file", "", "Path to a CSV file")
	flag.StringVar(&col, "column", "1", "Numerical columns")
	flag.StringVar(&datecol, "date-column", "", "The date column")
	flag.StringVar(&datefmt, "date-format", "", "The date format")
	flag.StringVar(&dateagg, "date-aggregation", "", "Aggregate summary stats by day, month, year")

	flag.Parse()

	options := BuildOptions(
		file,
		col,
		datecol,
		datefmt,
		dateagg,
	)

	summary := Summarize(options)

	for _, v := range summary.sort() {
		fmt.Printf("%s\n", v.name)
		fmt.Printf("\tsum\t%v\n", v.sum)
		fmt.Printf("\tavg\t%v\n", v.average)
		fmt.Printf("\tmin\t%v\n", v.min)
		fmt.Printf("\tmax\t%v\n", v.max)
		fmt.Printf("\tmed\t%v\n", v.median)
	}
}
