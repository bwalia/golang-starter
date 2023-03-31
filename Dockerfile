FROM golang:alpine3.17

RUN apk add bash
RUN apk add curl

WORKDIR /

COPY / /

ENTRYPOINT [ "bash /restapi-test-unit-test-1/test.sh" ]