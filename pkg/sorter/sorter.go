package sorter

import (
	"sort"

	"github.com/JoaoDanielRufino/gcloc/pkg/scanner"
)

type Result struct {
	Name       string
	Lines      int
	CodeLines  int
	BlankLines int
	Comments   int
}

type SortedSummary struct {
	Results         []Result
	FilesByLanguage map[string]int
	TotalFiles      int
	TotalLines      int
	TotalCodeLines  int
	TotalBlankLines int
	TotalComments   int
}

type Sorter interface {
	OrderByLanguage(summary *scanner.Summary) *SortedSummary
	OrderByCodeLines(summary *scanner.Summary) *SortedSummary
	OrderByLines(summary *scanner.Summary) *SortedSummary
	OrderByComments(summary *scanner.Summary) *SortedSummary
	OrderByBlankLines(summary *scanner.Summary) *SortedSummary
}

type baseSorter struct {
	sortOrder string
}

func (b baseSorter) sortByCodeLines(results []Result) {
	if b.sortOrder == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].CodeLines
			b := results[j].CodeLines
			return a < b
		})
	} else if b.sortOrder == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].CodeLines
			b := results[j].CodeLines
			return a > b
		})
	}
}

func (b baseSorter) sortByLines(results []Result) {
	if b.sortOrder == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Lines
			b := results[j].Lines
			return a < b
		})
	} else if b.sortOrder == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Lines
			b := results[j].Lines
			return a > b
		})
	}
}

func (b baseSorter) sortByComments(results []Result) {
	if b.sortOrder == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Comments
			b := results[j].Comments
			return a < b
		})
	} else if b.sortOrder == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].Comments
			b := results[j].Comments
			return a > b
		})
	}
}

func (b baseSorter) sortByBlankLines(results []Result) {
	if b.sortOrder == "ASC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].BlankLines
			b := results[j].BlankLines
			return a < b
		})
	} else if b.sortOrder == "DESC" {
		sort.Slice(results, func(i, j int) bool {
			a := results[i].BlankLines
			b := results[j].BlankLines
			return a > b
		})
	}
}
