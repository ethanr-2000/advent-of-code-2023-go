#!/bin/bash

build() {
	go build main.go
}

run() {
	if ! ./main > /dev/null 2>&1; then
		echo "Error in $1"
		exit 1
	fi
}

for year_folder in [0-9][0-9][0-9][0-9]/; do
	for day_folder in "$year_folder"day[0-9][0-9]/; do
		if [ -d "$day_folder" ]; then
			cd $day_folder
			build
			duration=$({ time (run $day_folder) 2>&1; } | grep ^real | awk '{print $2}') && echo "$duration for $day_folder"
			cd ../..
		fi
	done
done

wait

echo "All parts ran successfully"
