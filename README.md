# 🎄 advent-of-code 🎄

![Go](https://github.com/jcockbain/advent-of-code/workflows/Go/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/jcockbain/advent-of-code)

Solutions to 2021 and 2022 Advent of code. 

## Summary 

[Advent of Code](https://adventofcode.com/) is an annual advent-calendar of programming puzzles. Here are my solutions in Golang. 

## Running the Code

To fetch the input, and create a template dir for each day: 

```shell
./new_day {2022} {day1}
```

To then run the solutions: 

```go
// using 2022 day1 as an example
cd 2022/day01

// run the binary
go run main.go

// run tests
go test

// benchmark solution
go test --bench=BenchmarkMain

```

## Runtimes

The results are found using a `BenchmarkMain` benchmark in each solution. This table is generated by running the `benchmark.sh` script.

| Totals | 10.4s  | 12.8ms |
|  DAY   |  2021  |  2022  |
|--------|--------|--------|
|      1 | 89.8µs | 233µs  |
|      2 | 427µs  | 568µs  |
|      3 | 399µs  | 807µs  |
|      4 | 9.21ms | 853µs  |
|      5 | 26.2ms | 445µs  |
|      6 | 59.4µs | 334µs  |
|      7 | 29.5ms | 502µs  |
|      8 | 2.16ms | 7.24ms |
|      9 | 5.26ms | 1.9ms  |
|     10 | 29.6µs | 1.08ms |
|     11 | 4.8ms  |
|     12 | 280ms  | 268ms  |
|     13 | 1.31ms |
|     14 | 744µs  |
|     15 | 1.36s  |
|     16 | 239µs  |
|     17 | 2.05ms |
|     18 | 357ms  |
|     19 | 4.43s  |
|     20 | 440ms  |
|     21 | 728ms  |
|     22 | 4.42ms |
|     23 | 1.19s  |
|     24 | 15.2µs |
|     25 | 1.56s  |
| Totals | 10.4s  | 281ms  |
