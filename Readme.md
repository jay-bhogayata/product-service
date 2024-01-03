# 🚀 Product Service

Product Service is a robust microservice designed to manage products. It's built with Go and can be easily run with Make.

## 📚 Description

This service is part of a microservices architecture system. It's responsible for all operations related to products, such as creating, updating, deleting, and retrieving product data.

## 🛠️ Requirements

To run this service, you'll need:

- [Go](https://golang.org/dl/) 1.21 or later
- [Make](https://www.gnu.org/software/make/)

## 🚀 Getting Started

You can run the service in two ways:

1. Using Make:

    ```sh
    make run
    ```

2. Directly with Go:

    ```sh
    go run main.go
    ```

## 📖 Documentation

Once the service is running, you can access the API documentation at the following URL:

[Swagger API Documentation](http://localhost:8080/swagger/index.html)

## 💻 Local Development

For local development, we recommend using [Air](https://github.com/cosmtrek/air) for hot reloading. This will automatically rebuild and restart your service whenever you make changes to the source code.

To start the service with Air, run:

```sh
air
```
note : air does hot reload swagger docs, so you need to restart the service to see the changes.