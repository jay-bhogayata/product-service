[![ci workflow](https://github.com/jay-bhogayata/product-service/actions/workflows/ci.yaml/badge.svg?branch=main)](https://github.com/jay-bhogayata/product-service/actions/workflows/ci.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/jay-bhogayata/product-service)](https://goreportcard.com/report/github.com/jay-bhogayata/product-service)

# 🚀 Product Service

Welcome to Product Service, a robust and efficient microservice designed to manage products. Built with the power of Go.
## 📘 Description

Product Service is the backbone of all operations related to products. It handles creating, updating, deleting, and retrieving product data, ensuring smooth and efficient product management.

## 🛠️ Prerequisites

Before you begin, ensure you have met the following requirements:

- You have installed [Go](https://golang.org/dl/) 1.21 or later.
- You have installed [Make](https://www.gnu.org/software/make/).

## ⚙️ Configuration

This service relies on a configuration file for operation. Follow these steps to set it up:

1. Locate the `sample.config.yaml` file in the root directory of the project.
2. Rename `sample.config.yaml` to `cfg.yaml`.
3. Update the variables in `cfg.yaml` with your own values.

## 🚀 Quickstart

Kickstart the service in one of two ways:

1. Using Make:

    ```sh
    make run
    ```

2. Directly with Go:

    ```sh
    go run main.go
    ```

## 📖 API Documentation

Once the service is up and running, you can access the API documentation at:

[Swagger API Documentation](http://localhost:8080/swagger/index.html)

## 💻 Local Development

For a seamless local development experience, we recommend using [Air](https://github.com/cosmtrek/air) for hot reloading. Air automatically rebuilds and restarts your service whenever you make changes to the source code.

To start the service with Air, run:

```sh
air
```