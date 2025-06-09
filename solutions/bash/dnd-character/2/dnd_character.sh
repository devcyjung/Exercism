#!/usr/bin/env bash

diceroll() {
    shuf -i 1-6 -n 1
}

ability() {
    local mini=7
    local sum=0
    for _ in {1..4}; do
        local roll
        roll=$(diceroll)
        sum=$((sum + roll))
        if ((roll < mini)); then
            mini=$roll
        fi
    done
    echo $((sum - mini))
}

modifier() {
    local v=$(($1 - 10))
    if ((v < 0)); then
        v=$((v + v % 2))
    else
        v=$((v - v % 2))
    fi
    v=$((v / 2))
    echo $v
}

generate() {
    local str
    str=$(ability)
    local dex
    dex=$(ability)
    local con
    con=$(ability)
    local int
    int=$(ability)
    local wis
    wis=$(ability)
    local chr
    chr=$(ability)
    local mod
    mod=$(modifier "$con")
    local hit
    hit=$((10 + mod))
    printf "strength %s\n" "$str"
    printf "dexterity %s\n" "$dex"
    printf "constitution %s\n" "$con"
    printf "intelligence %s\n" "$int"
    printf "wisdom %s\n" "$wis"
    printf "charisma %s\n" "$chr"
    printf "hitpoints %s\n" "$hit"
}

main() {
    local cmd=$1
    if [[ "$cmd" = "modifier" ]]; then
        modifier "$2"
    elif [[ "$cmd" = "generate" ]]; then
        generate
    fi
}

main "$@"
