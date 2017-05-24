// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/nagai1110/Go/ch04/ex10/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Printf("%d issues:\n", result.TotalCount)
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// }

	now := time.Now()
	lastMonth := now.AddDate(0, -1, 0)
	lastYear := now.AddDate(-1, 0, 0)

	fmt.Println("Since last maonth")
	filter(result, &lastMonth, &now)
	fmt.Println("Since last year")
	filter(result, &lastYear, &lastMonth)
	fmt.Println("Others")
	filter(result, nil, &lastYear)
}

func filter(result *github.IssuesSearchResult, startDate *time.Time, endDate *time.Time) {
	for _, issue := range result.Items {
		date := issue.CreatedAt

		if (startDate == nil || (*startDate).Equal(date) || (*startDate).After(date)) &&
			(endDate == nil || (*endDate).Equal(date) || (*endDate).Before(date)) {
			fmt.Printf("#%-5d %v %9.9s %.55s\n", issue.Number, issue.CreatedAt, issue.User.Login, issue.Title)
		}
	}
}
