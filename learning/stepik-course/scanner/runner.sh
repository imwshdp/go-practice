#!/bin/bash

if [ ! -f "$1" ]; then
  echo "File $1 is not existed"
  exit 1
fi

cat "$1" | go run main.go
