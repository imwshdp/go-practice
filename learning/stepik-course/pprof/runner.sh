#!/bin/bash

# =====
# check results
# =====
go test -bench . -benchmem

# =====
# fast all info
# =====
# go test -bench=. -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof -memprofilerate=1

# =====
# cpu/mem benchmarks
# =====
# go test -bench=. -cpuprofile=cpu.prof -memprofile=mem.prof -memprofilerate=1
# go tool pprof cpu.prof
# go tool pprof mem.prof

# =====
# pprof debug
# =====
# top
# alloc_space
# list SlowSearch
# list FastSearch