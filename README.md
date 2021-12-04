# 🎄 advent-of-code-2021 🎄

![Go](https://github.com/jcockbain/advent-of-code-2021/workflows/Go/badge.svg)

Solutions to 2021 Advent of Code. 

## Summary 

[Advent of Code](https://adventofcode.com/) is an annual advent-calendar of programming puzzles. Here are my 2021 solutions in Golang. 

## Running the Code

To fetch the input, and create a template dir for each day: 

```shell
./new_day 2021 {day1}
```

To then run the solutions: 

```go
cd day01 

// defaults to input.txt in current dir
go run main.go

// run with custom input (I normally use the test input given with the problem)
go run main.go test1.txt

```