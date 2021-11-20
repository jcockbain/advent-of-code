#! /bin/bash

YEAR=$1
DAY=$2

echo "Creating Boilerplate for year ${YEAR} day ${DAY}"

# STRLENGTH=$(echo -n $DAY | wc -m)
# echo $STRLENGTH
if [[ ${#DAY} -ge 1 ]]
then
    DISPLAY_DAY="0${DAY}"
else
    DISPLAY_DAY=$DAY
fi

echo $DISPLAY_DAY

GOROOT=day${DISPLAY_DAY}
cp -n -R template $GOROOT

# copy input

INPUT_URL="https://adventofcode.com/2020/day/1/input"
TEMP_INPUT="temp-input.txt"

# "session cookie" must be given as input is per-user...
# this can be grabbed from the URL of any input page online e.g https://adventofcode.com/2019/day/2/input
# I have this set as an environment variable here

curl "${INPUT_URL}" -H "cookie: session=${AOC_COOKIE}" -o "${TEMP_INPUT}" 2>/dev/null

cp ${TEMP_INPUT} ${GOROOT}/input.txt

# clean up

rm ${TEMP_INPUT}