

# go-product-service-grpc

This repository contains a gRPC service and client implementation for a product service in Go.

## Client

The client implementation can be found in the `client` directory. It communicates with the gRPC server to retrieve product information. The client code is located in the `client.go` file.

To run the client, use the following command:

```bash
go run client/client.go
```

Server
------

The server implementation can be found in the `server` directory. It provides a gRPC service for retrieving product information. The server code is located in the `server.go` file.

To run the server, use the following command:

`go run server.go`

Protobuf
--------

The protocol buffer definition file (`product.proto`) can be found in the root directory of this repository. It defines the `Product` message and the `ProductService` service.
