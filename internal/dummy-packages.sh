#!/usr/bin/env bash

set -euo pipefail

cd "$(dirname "$0")"

for d in $(find libusb hidapi -type d)
do
	echo "package $(basename $d)" >"$d/pkg.go"
done

pkg=$(cd ..; go list)/internal

cksum="$(find hidapi libusb -type f -name '*.[ch]' | sort | xargs cat | openssl sha256)"

{
	echo 'package internal'
	echo
	echo 'import ('
	for d in $(find hidapi libusb -type d);
	do
		printf '\t_ "%s/%s"\n' "$pkg" "$d"
	done
	echo ')'
	echo
	echo '// Checksum is a checksum of *.c and *.h files in order to trigger a rebuild'
	echo '// of package ./.. if one of those files changes.'
	echo '// See https://github.com/golang/go/issues/26366#issuecomment-405683150'
	printf 'const Checksum = "%s"\n' "$cksum"
} >pkg.go
