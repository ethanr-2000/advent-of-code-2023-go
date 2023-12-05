#!/bin/bash

run_tests() {
	if ! go test > /dev/null 2>&1; then
		echo "Tests failed in $1"
		exit 1
	fi
}

for year_folder in [0-9][0-9][0-9][0-9]/; do
	for day_folder in "$year_folder"day[0-9][0-9]/; do
		if [ -d "$day_folder" ]; then
			cd $day_folder
			run_tests $day_folder &
			cd ../..
		fi
	done
done

for pkg_folder in pkg/*; do
	if [ -d "$pkg_folder" ]; then
		cd $pkg_folder
		run_tests $pkg_folder &
		cd ../..
	fi
done

wait

echo "All tests ran successfully"
