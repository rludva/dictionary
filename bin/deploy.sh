#!/bin/bash

# Load common dependencies for bash scripts..
MY_PATH=$(dirname "$0")

OUTPUT_DIRECTORY="$MY_PATH/../output"
if [ ! -e "$OUTPUT_DIRECTORY/dictionary_linux_amd64" ]; then
  echo "Nothing builded.."
  exit
fi

# Easy copy..
cp "$OUTPUT_DIRECTORY/dictionary_linux_amd64" $HOME/.local/bin/dictionary
