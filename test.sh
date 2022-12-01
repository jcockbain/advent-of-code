#!/bin/bash

y=${1}

d=${2:-$DAY}
day=day${d}
day_path=${y}/${day}

echo $day_path

if [ -z ${d} ]
then
  echo "Set \$DAY or pass day directory as an arg to test single file"
  # test all
  go test ./${y}/...
elif [ ! -d ${day_path} ]
then
  echo "$day_path is not a directory!"
  exit 1
else
  # test individual day
  cd ${day_path} && go test
fi
