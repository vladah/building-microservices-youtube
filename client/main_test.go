package main

import (
	"testing"

	"github.com/vladah/building-microservices-youtube/client/client"
)

func TestOutClient(t *testing.T) {
	c := client.Default
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)
}
