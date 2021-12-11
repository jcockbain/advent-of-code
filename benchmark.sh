#!/bin/bash

for i in $(seq -f "%02g" 1 25)
do  
    if [ -d "day${i}" ]
    then
        echo "benchmarking day ${i}"
        cd "day${i}" && go test --bench=BenchmarkMain | grep "BenchmarkMain-" | awk '{print $3}' > ../benchmarks/results/${i}.txt
        cd ..
    fi
done

cd benchmarks || exit
go run benchmarks.go
cd ..
