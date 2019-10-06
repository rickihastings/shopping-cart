# Shopping Cart

This repo implements a basic shopping cart microservice in Go with a memory backend. The following endpoints are included

- Adding Item
- Removing Item
- Clearing Items
- Listing Items in basket
- Calculating delivery

## Installation

1. Install Go first. I like to use the official instructions, there is installation packages for OS X and Windows. https://golang.org/doc/install
2. Once installed, test go is working. `go version`. You should see something like: `go version go1.13 darwin/amd64`
3. Install this package. Go has a clever package manager, no need to git clone packages, we can download it using the go tools. `go get rickihastings/shopping-cart`

## Running

`go run ./cmd/shopping-cart/`

## Testing

`go test ./pkg/api`

## Deployment

If this was to be deployed in the real world. The application would be packaged into a Docker container, and deployed to Kubernetes with a horizontal pod autoscaler to ensure it can scale appropriately. Builds and deployments could be managed with an AWS codepipeline and AWS code deploy scripts.

As the data is currently just JSON and non-relational, it might make sense to use a NoSQL database such as MongoDB, or DynamoDB. This could be integrated quite easily by modifying the `store.go` file.

## Uninstall

If have never installed Go before, and wish to uninstall Go to keep your system clean, you can use the following guide. https://golang.org/doc/install#uninstall
