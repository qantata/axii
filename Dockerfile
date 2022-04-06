FROM golang:1.17-alpine as build

WORKDIR /usr/src/app

COPY . .

RUN go build .

FROM alpine

RUN adduser -D appuser
USER appuser

COPY --from=build /usr/src/app/axii ./axii

CMD ./axii