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

|  DAY   |  2021  |  2022  |
|--------|--------|--------|
|      1 | 89.8µs | 233µs  |
|      2 | 424µs  | 568µs  |
|      3 | 399µs  | 821µs  |
|      4 | 9.21ms | 855µs  |
|      5 | 447µs  | 26.2ms |
|      6 | 59.4µs | 333µs  |
|      7 | 504µs  | 29.5ms |
|      8 | 2.16ms | 7.13ms |
|      9 | 5.26ms | 1.9ms  |
|     10 | 1.08ms | 29µs   |
|     11 | 4.8ms  | 19.9µs |
|     12 | 280ms  | 247ms  |
|     13 | 1.31ms |
|     14 | 744µs  | 86.3ms |
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
| Totals | 10.4s  | 346ms  |
