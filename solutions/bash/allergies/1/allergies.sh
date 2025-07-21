#!/usr/bin/env bash

score="$1"
flag="$2"

declare -A allergens
allergens[eggs]=1
allergens[peanuts]=2
allergens[shellfish]=4
allergens[strawberries]=8
allergens[tomatoes]=16
allergens[chocolate]=32
allergens[pollen]=64
allergens[cats]=128

declare -a allergens_sorted
allergens_sorted=(
    "eggs" "peanuts" "shellfish" "strawberries" "tomatoes" "chocolate" "pollen" "cats"
)

allergic_to() {
	if [ $((allergens[$1] & $2)) -gt 0 ]; then
		echo "true"
	else
		echo "false"
	fi
}

if [ "$flag" = "allergic_to" ]; then
	food="$3"
	allergic_to "$food" "$score"
elif [ "$flag" = "list" ]; then
	declare -a buffer=()
	for food in "${allergens_sorted[@]}"; do
		result="$(allergic_to "$food" "$score")"
		if [ "$result" = "true" ]; then
			buffer+=("$food")
		fi
	done
	echo "${buffer[*]}"
fi
