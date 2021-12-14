package main

import (
	"embed"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

//go:embed results
var resultsFS embed.FS

type benchmarkData map[int]time.Duration

func main() {
	benchmarks, err := loadBenchmarks()
	if err != nil {
		panic(err)
	}
	table := getTable(benchmarks)
	fmt.Print(table)
	ioutil.WriteFile("README.md", []byte(table), 0644)
}

func getTable(benchmarks benchmarkData) string {
	data := [][]string{}

	totalRuntime := time.Duration(0)
	for day, runtime := range benchmarks {
		data = append(data, []string{fmt.Sprint(day), formatDuration(runtime)})
		totalRuntime += runtime
	}

	sort.SliceStable(data, func(i, j int) bool {
		return toInt(data[i][0]) < toInt(data[j][0])
	})
	data = append(data, []string{"Total", formatDuration(totalRuntime)})

	tableString := &strings.Builder{}
	table := tablewriter.NewWriter(tableString)
	table.SetHeader([]string{"Day", "Runtime"})
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	return tableString.String()
}

func loadBenchmarks() (benchmarkData, error) {
	benchmarks := benchmarkData{}

	dayResults, err := resultsFS.ReadDir("results")
	if err != nil {
		return nil, err
	}
	for _, dayResult := range dayResults {
		day, err := strconv.Atoi(strings.TrimSuffix(dayResult.Name(), ".txt"))
		if err != nil {
			return nil, err
		}
		path := fmt.Sprintf("results/%s", dayResult.Name())
		result, err := resultsFS.ReadFile(path)
		if err != nil {
			return nil, err
		}
		benchmark, err := strconv.ParseFloat(strings.TrimSpace(string(result)), 64)
		if err == nil {
			benchmarks[day] = time.Duration(int64(math.Round(benchmark)))
		}
	}
	return benchmarks, nil
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
