FROM golang:alpine AS builder

# Build the binary
WORKDIR /src

RUN mkdir app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./app

WORKDIR /src/app

RUN go build -o user-service

# Serve the app
FROM alpine

WORKDIR /app

COPY --from=builder /src/app/user-service /app/

EXPOSE 8080

ENTRYPOINT [ "/app/user-service" ]