#!/bin/bash

OUTPUT_DIRECTORY="./output"
if [ ! -e "$OUTPUT_DIRECTORY" ]; then
  mkdir $OUTPUT_DIRECTORY
fi

function build() {
	GOOS="$1"
	GOARCH="$2"

	echo "Building for $GOOS/$GOARCH.."

	OUTPUT="$OUTPUT_DIRECTORY/dictionary_${GOOS}_${GOARCH}"
  if [ -e "$OUTPUT" ]; then
		rm "$OUTPUT"
  fi

	env GOOS="$GOOS" GOARCH="$GOARCH" go build -o "$OUTPUT"
}

build linux amd64
#build windows amd64
#build darwin amd64 
ls -l "$OUTPUT_DIRECTORY"
