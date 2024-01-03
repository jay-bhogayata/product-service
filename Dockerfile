FROM golang:1.21.5-alpine3.18x AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV GOCACHE=/root/.cache/go-build
RUN --mount=type=cache,target="/root/.cache/go-build" CGO_ENABLED=0 GOOS=linux go build -o product-api 

FROM gcr.io/distroless/static-debian12
WORKDIR /
COPY --from=builder /app/product-api /product-api
EXPOSE 8080
CMD [ "./product-api" ]