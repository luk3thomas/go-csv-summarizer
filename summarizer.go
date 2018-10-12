package main

import (
	"io"
	"log"
	"math"
	"sort"
	"strconv"
	"time"
)

// Result summarized stats
type Result struct {
	name    string
	sum     float64
	min     float64
	max     float64
	average float64
	median  float64
	count   int64
}

// ResultSet a set of results
type ResultSet map[time.Time]Result

type list []float64
type aggregatedMap map[time.Time][]float64

func (r ResultSet) sort() []Result {
	var keys []time.Time
	var result []Result
	for t := range r {
		keys = append(keys, t)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	for _, k := range keys {
		result = append(result, r[k])
	}
	return result
}

// Summarize a list of list
func Summarize(opts Options) ResultSet {
	if opts.datecol == 0 {
		list := summarizeList(opts)
		list.name = "all"
		res := make(ResultSet)
		res[time.Now()] = list
		return res
	}
	return aggList(opts)
}

func aggList(opts Options) ResultSet {
	var result = make(ResultSet)
	var list = make(aggregatedMap)
	for {
		record, err := opts.file.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		date := getColumn(opts.datecol, record)
		value := getColumnValue(opts.column, record)
		if value != nil && date != nil {
			list = appendValue(list, opts, *date, *value)
		}
	}

	for k, v := range list {
		r := calc(v)
		r.name = k.Format("2006-01")
		result[k] = r
	}

	return result
}

func appendValue(r aggregatedMap, opts Options, date string, value float64) aggregatedMap {
	var aggTime time.Time
	t, err := time.Parse(opts.datefmt, date)
	if err != nil {
		return r
	}

	if opts.dateagg == DateFormatYear {
		aggTime, _ = time.Parse("2006", t.Format("2006"))
	} else {
		aggTime, _ = time.Parse("2006-01", t.Format("2006-01"))
	}

	val, hasKey := r[aggTime]

	if hasKey {
		l := append(val, value)
		r[aggTime] = l
	} else {
		r[aggTime] = list{value}
	}
	return r
}

func summarizeList(opts Options) Result {
	var n list
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

func calc(n list) Result {
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
	r.median = getMedian(n, r.count)
	return r
}

func getMedian(n list, count int64) float64 {
	index := float64(count / 2)
	middle := int(math.Floor(index))
	sort.Float64s(n)
	return n[middle]
}

func getColumn(col Column, record []string) *string {
	var index = int(col - 1)
	if index > len(record) {
		return nil
	}
	return &record[index]
}

func getColumnValue(col Column, record []string) *float64 {
	v := getColumn(col, record)
	value, err := strconv.ParseFloat(*v, 64)
	if err != nil {
		return nil
	}
	return &value
}
