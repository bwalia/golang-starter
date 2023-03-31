#!/bin/bash

set -x

#go run main.go
go build -o /resapi-smoke-test

/usr/bin/supervisord -c /etc/supervisord.conf

echo "REST API Golang unit test suit"

sleep 10

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
