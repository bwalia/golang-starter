#!/usr/bin/env sh

set -e

echo "=== Test GET ==="
#curl -X GET -s localhost:8080 | rg -q "get" 
#curl -X POST -s localhost:8080 | rg -q "post"
RESP=$(curl -X GET -s localhost:8080)
if [[ "$RESP" = "get" ]] ; then
echo "Get API QA Passed: Success"
else
echo "Get API QA Failed: Error"
fi

echo "=== Test POST ==="

RESP=$(curl -X POST -s localhost:8080)
if [[ "$RESP" = "post" ]] ; then
echo "Post API QA Passed: Success"
else
echo "Post API QA Failed: Error"
fi
