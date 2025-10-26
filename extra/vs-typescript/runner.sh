#!/bin/bash

echo "=== Golang ==="
go run go.go &
GO_PID=$!

npx -p typescript tsc ts.ts && node ts.js &
TS_PID=$!

wait $GO_PID

echo -e "\n\n=== JavaScript ==="
wait $TS_PID
