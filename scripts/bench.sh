#!/usr/bin/env bash

set -e

ab \
  -n 50000 \
  -c 400 \
  -k -v 1 \
  -H "Accept-Encoding: gzip, deflate" \
  -T "application/json" \
  -p scripts/bench.txt "http://localhost:8080/api/v1" > .results

  ab \
  -n 50000 \
  -c 400 \
  -k -v 1 \
  -H "Accept-Encoding: gzip, deflate" \
  -T "application/json" \
  "http://localhost:8080/api/v1?key=ok" > .results
