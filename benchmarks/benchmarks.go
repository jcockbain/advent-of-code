package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"math"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

var years = []int{2021, 2022}

//go:embed results/*
var resultsFS embed.FS

// year to duration
type rowData map[string]time.Duration

// day to years e.g 1 -> {2021: 1s, 2022: 2s}
type tableData map[int]rowData

var tData tableData

func main() {
	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	headers := []string{"Day"}

	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")

	tData = make(tableData)

	totalRuntimes := map[int]time.Duration{}
	for _, y := range years {
		yearString := strconv.FormatInt(int64(y), 10)
		headers = append(headers, yearString)
		totalRuntime, err := populateYearData(yearString)
		if err != nil {
			panic(err)
		}
		totalRuntimes[y] = totalRuntime
	}
	table.SetHeader(headers)
	for i := 1; i <= 25; i++ {
		day := i
		data := tData[i]
		dayString := strconv.FormatInt(int64(day), 10)
		row := []string{dayString}
		for _, t := range data {
			row = append(row, formatDuration(t))
		}
		table.Append(row)
	}
	footerRow := []string{"Totals"}
	for _, year := range years {
		total := totalRuntimes[year]
		footerRow = append(footerRow, formatDuration(total))
	}
	table.Append(footerRow)
	table.Render()
	ioutil.WriteFile("README.md", []byte(tableString.String()), 0644)
	fmt.Print(tableString.String())
}

func populateYearData(year string) (time.Duration, error) {
	dayResults, err := resultsFS.ReadDir(path.Join("results", year))
	if err != nil {
		return time.Duration(0), err
	}
	totalDuration := time.Duration(0)
	for _, dayResult := range dayResults {
		day, err := strconv.Atoi(strings.TrimSuffix(dayResult.Name(), ".txt"))
		if err != nil {
			return time.Duration(0), err
		}
		path := fmt.Sprintf("results/%s/%s", year, dayResult.Name())
		result, err := resultsFS.ReadFile(path)
		if err != nil {
			return time.Duration(0), err
		}
		benchmark, err := strconv.ParseFloat(strings.TrimSpace(string(result)), 64)
		if err == nil {
			if tData[day] == nil {
				tData[day] = make(rowData)
			}
			duration := time.Duration(int64(math.Round(benchmark)))
			totalDuration += duration
			tData[day][year] = duration
		}
	}
	return totalDuration, nil
}

func formatDuration(dur time.Duration) string {
	for interval := time.Nanosecond; interval <= time.Second; interval *= 10 {
		if dur >= 100*interval {
			dur = dur.Round(interval)
		}
	}
	return dur.String()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}
