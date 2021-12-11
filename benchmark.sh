#!/bin/bash

for i in $(seq -f "%02g" 1 25)
do  
    echo "benchmarking day ${i}"
    cd "day${i}" && go test --bench=BenchmarkMain | grep "BenchmarkMain-" | awk '{print $3}' > ../benchmarks/results/${i}.txt
    cd ..
done
