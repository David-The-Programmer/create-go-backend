# syntax=docker/dockerfile:1

ARG GO_VERSION

FROM golang:${GO_VERSION}

WORKDIR /backend

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o ./run

EXPOSE 3000

CMD ["./run"]
