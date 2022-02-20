#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")"

for d in $(find libusb hidapi -type d)
do
	echo "package $(basename $d)" >"$d/pkg.go"
done

pkg=$(cd ..; go list)/internal

{
	echo 'package internal'
	echo
	echo 'import ('
	for d in $(find hidapi libusb -type d);
	do
		printf '\t_ "%s/%s"\n' "$pkg" "$d"
	done
	echo ')'
} >pkg.go
