# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /backend

COPY go.* ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

EXPOSE 3000

CMD ["/backend"]
