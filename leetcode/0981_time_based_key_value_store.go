package main

import "fmt"

func main() {
	tm := Constructor()
	tm.Set("foo", "bar", 1)
	fmt.Println("foo, 1", tm.Get("foo", 1))
	fmt.Println("foo, 3", tm.Get("foo", 3))
	tm.Set("foo", "bar2", 4)
	fmt.Println("foo, 4", tm.Get("foo", 4))
	fmt.Println("foo, 5", tm.Get("foo", 5))
	// fmt.Println("foo, 0", tm.Get("foo", 0))
}

type Val struct {
	timestamp int
	value     string
}

type TimeMap map[string][]Val

func Constructor() TimeMap {
	return TimeMap{}
}

func (tm *TimeMap) Set(key, value string, timestamp int) {
	val := Val{timestamp: timestamp, value: value}
	(*tm)[key] = append((*tm)[key], val)
}

func (tm TimeMap) Get(key string, timestamp int) string {
	vals, ok := tm[key]
	if !ok || len(vals) == 0 {
		return ""
	}
	if vals[0].timestamp > timestamp {
		return ""
	}

	l := 0
	r := len(vals)
	m := (l + r) / 2
	for l < r {
		m = (l + r) / 2
		if vals[m].timestamp <= timestamp {
			l = m + 1
		} else {
			r = m
		}
	}
	return vals[m].value
}
