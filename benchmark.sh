#!/bin/bash

# example use
# ./benchmark.sh 2022 - runs all benchmarks for 2022 year and outputs table
# ./benchmark.sh 03 2022 - runs only benchmark for 2022 day 03

function bench {
    echo "benchmarking day $1 year $2"
    cd "$2/day$1" && go test --bench=BenchmarkMain | grep "BenchmarkMain-" | awk '{print $3}' > ../../benchmarks/results/$2/$1.txt
    cd ../..
}

if [ $# -eq 1 ]
then
    for i in $(seq -f "%02g" 1 25)
    do
        if [ -d "$1/day${i}" ]
        then
        bench $i $1
        fi
    done
else
    bench $1 $2
fi

cd benchmarks || exit
go run benchmarks.go
cd ..

# remove the old table from readme
sed -i '' '/|/d' README.md
# copy new table to readme
cat benchmarks/README.md >> README.md
