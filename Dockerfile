FROM golang:alpine3.17

RUN apk add bash
RUN apk add curl
RUN apk add supervisor

WORKDIR /

COPY /restapi-test-unit-test-1 /
COPY supervisord.conf /etc/

CMD [ "./test.sh" ]