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

## Uninstall

If have never installed Go before, and wish to uninstall Go to keep your system clean, you can use the following guide. https://golang.org/doc/install#uninstall
