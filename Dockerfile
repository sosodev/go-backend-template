# Using the Go image to make building the app easy
FROM golang:1.13 as builder

ADD . /app
WORKDIR /app

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -o backend

# Using the alpine image as a slim, but not minimal, runtime for the binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/backend .

EXPOSE 8080

CMD ["./backend"]

