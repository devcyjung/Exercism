#!/usr/bin/env bash

flag="$1"
p="$2"

powmod() {
	echo "($1 ^ $2) % $3" | bc
}

privateKey() {
	shuf -i 2-$(("$p" - 1)) -n 1
}

publicKey() {
	powmod "$g" "$private" "$p"
}

secret() {
	powmod "$public" "$private" "$p"
}

if [ "$flag" = "privateKey" ]; then
	privateKey "$p"
elif [ "$flag" = "publicKey" ]; then
	g="$3"
	private="$4"
	publicKey "$p" "$g" "$private"
elif [ "$flag" = "secret" ]; then
	public="$3"
	private="$4"
	secret "$p" "$public" "$private"
fi
