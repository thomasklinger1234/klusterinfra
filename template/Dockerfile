FROM golang:1.20.2-alpine3.17 as BUILDER

WORKDIR /opt/app

RUN apk add -U --no-cache build-base

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod verify

COPY . .

RUN go build -o bootstrap main.go

FROM alpine:3.17

COPY --from=BUILDER /opt/app/bootstrap /opt/app/bootstrap

USER daemon 

ENTRYPOINT [ "/opt/app/bootstrap" ]