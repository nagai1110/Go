// Copyright © 2017 Takaya.Nagai All Rights Reserved.
package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type lessFunc func(x, y *Track) bool

type customSort struct {
	t         []*Track
	lessFuncs []lessFunc
}

func byTitle(x, y *Track) bool  { return x.Title < y.Title }
func byArtist(x, y *Track) bool { return x.Artist < y.Artist }
func byAlbum(x, y *Track) bool  { return x.Album < y.Album }
func byYear(x, y *Track) bool   { return x.Year < y.Year }
func byLength(x, y *Track) bool { return x.Length < y.Length }

func (x customSort) Len() int { return len(x.t) }

func (x customSort) Less(i, j int) bool {
	ti, tj := x.t[i], x.t[j]

	for _, lessFunc := range x.lessFuncs {
		if lessFunc(ti, tj) {
			return true
		} else if lessFunc(tj, ti) {
			return false
		}
	}

	return false
}

func (x customSort) Swap(i, j int) { x.t[i], x.t[j] = x.t[j], x.t[i] }

func main() {
	fmt.Println("=== sort by Title, Length ==")
	cs := customSort{t: tracks, lessFuncs: []lessFunc{byTitle, byLength}}
	sort.Sort(&cs)
	printTracks(tracks)
}
