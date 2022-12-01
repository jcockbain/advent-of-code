#!/bin/bash

YEAR=$1
DAY=$2

echo "Creating Boilerplate for year ${YEAR} day ${DAY}"
day_length=${#DAY}

if [ $day_length == 1 ]
then DISPLAY_DAY="0${DAY}"
else DISPLAY_DAY=$DAY
fi

GOROOT=${YEAR}/day${DISPLAY_DAY}
cp -n -R dayxx $GOROOT

INPUT_URL="https://adventofcode.com/${YEAR}/day/${DAY}/input"
TEMP_INPUT="temp-input.txt"

# "session cookie" must be given as input is per-user...
# this can be grabbed from the URL of any input page online e.g https://adventofcode.com/2019/day/2/input

curl "${INPUT_URL}" -H "cookie: session=${AOC_COOKIE}" -o "${TEMP_INPUT}" 2>/dev/null
cp ${TEMP_INPUT} ${GOROOT}/input.txt
rm ${TEMP_INPUT}
