FROM golang:alpine3.17

RUN apk add bash
RUN apk add curl

WORKDIR /opt

COPY / /opt

ENTRYPOINT [ "/opt/restapi-test-unit-test-1/test.sh" ]