FROM golang:1.22-alpine3.20 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o sample

FROM alpine:3.20.3 AS production
RUN apk --no-cache add ca-certificates
WORKDIR /app

COPY --from=builder /app/sample ./

EXPOSE 8082
ENTRYPOINT [ "./sample" ]
