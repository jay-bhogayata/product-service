FROM golang:1.21.5-alpine3.18 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o product-servive

FROM alpine:3.14.2
WORKDIR /
COPY --from=builder /app/product-servive .   
EXPOSE 8080
CMD ["./product-servive"]
