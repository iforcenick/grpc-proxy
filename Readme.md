# grpc-proxy

## 👷 Background

I got this R&D Recruiting Challenge from Lavanet company.

## ✨ Features

1. The gRPC Server listens to requests from the user (you can send the
    requests using grpcurl ) and uses its implemented client to forward the
    request to the Osmosis public RPC endpoint all endpoints are located here:

    https://docs.osmosis.zone/developing/network/public-endpoints.html#official-
    endpoints
    
    Implemented support for at least `cosmos.base.tendermint.v1beta1.Service` service.


2. Edesign a state tracker that will query
    your server for the latest block information using the
    `cosmos.base.tendermint.v1beta1.Service.GetLatestBlock` API.

> ⚠️ WARNING: This project can only be run on Golang version greather than 1.16.0

## 🧱 Prerequisites

1. [Go](https://nodejs.org/) `>=1.18.0` (might be able to run on lower versions as well, not tested)

## 🎉 Getting started

```bash
# clone from github
git clone https://github.com/iforcenick/grpc-proxy

# cd into repo directory
cd grpc-proxy

# Install dependencies for server app
cd server
go mod tidy
cd ..

# Install dependencies for client app
cd client
go mod tidy
cd ..

# Run server
go run server

# Run client
go run client

# Test server endpoints
go test server
```

## 🚚 Installation

1. Git clone this repo somewhere you like it
2. `go mod tidy` inside the server directory and the client directory
3. Run gRPC server by executing `go run server` inside the root directory
4. `go run client` inside the root directory to run status tracker client
5. Run `go test server` to test the proxy app

## 🚧 Roadmap

- [ ] Support all services and not only `cosmos.base.tendermint.v1beta1.Service`
- [ ] The gRPC Services you Support, will be added automatically using a script instead of manually adding them.
- [ ] The gRPC server will be built dynamically by requesting the client what protobufs the endpoint supports

## 📄 License

This project is licensed under the [**MIT License**](LICENSE).
