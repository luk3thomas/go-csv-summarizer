package main

import (
	"io"
	"log"
	"strconv"
)

// Result summarized stats
type Result struct {
	sum     float64
	min     float64
	max     float64
	average float64
	count   int64
}

type numbers []float64

// Summarize a list of numbers
func Summarize(opts Options) Result {
	var n numbers
	for {
		record, err := opts.file.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		value := getColumnValue(opts.column, record)
		if value != nil {
			n = append(n, *value)
		}
	}
	return calc(n)
}

func calc(n numbers) Result {
	r := Result{}
	r.count = int64(len(n))
	for i, e := range n {
		if i == 0 {
			r.min = e
			r.max = e
		}
		if e < r.min {
			r.min = e
		}
		if e > r.max {
			r.max = e
		}
		r.sum = r.sum + e
	}
	r.average = r.sum / float64(r.count)
	return r
}

func getColumnValue(col Column, record []string) *float64 {
	var index = int(col - 1)
	if index > len(record) {
		return nil
	}
	value, err := strconv.ParseFloat(record[index], 64)
	if err != nil {
		return nil
	}
	return &value
}
