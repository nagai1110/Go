// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": false},
	"calculus":   {"linear algebra": false},

	"compilers": {
		"data structures":       false,
		"formal languages":      false,
		"computer organization": false,
	},

	"data structures":       {"discrete math": false},
	"databases":             {"data structures": false},
	"discrete math":         {"intro to programming": false},
	"formal languages":      {"discrete math": false},
	"networks":              {"operating systems": false},
	"operating systems":     {"data structures": false, "computer organization": false},
	"programming languages": {"data structures": false, "computer organization": false},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

// TODO:最初のソート削除
func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}

	visitAll(keys)
	return order
}
