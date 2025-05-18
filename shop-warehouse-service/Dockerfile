FROM golang:alpine AS builder

# Build the binary
WORKDIR /src

RUN mkdir app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . ./app

WORKDIR /src/app

RUN go build -o shop-warehouse-service

# Serve the app
FROM alpine

WORKDIR /app

COPY --from=builder /src/app/shop-warehouse-service /app/

EXPOSE 8082

ENTRYPOINT [ "/app/shop-warehouse-service" ]