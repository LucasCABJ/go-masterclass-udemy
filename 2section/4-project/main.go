package main

import (
	"fmt"
	"strings"
)

var productPrices = map[string]float64 {
	"tshirt": 20.0,
	"mug": 12.50,
	"hat": 18.0,
	"book": 25.99,
}

func calculateItemPrice(product string) (float64, bool) {
	price, found := productPrices[product] 
	if !found && strings.HasSuffix(product, "_SALE") {
		originalItemCode := strings.TrimSuffix(product, "_SALE")
		if basePrice, found := productPrices[originalItemCode]; found { 
			salePrice := basePrice * 0.9
			fmt.Printf(" - Item %s Sale! Original: %.2f, Sale Price: %.2f\n", originalItemCode, basePrice, salePrice)
			return salePrice, true
		}
	}
	return price, found
}

func main() {
	fmt.Println("-------------- Simple Sales Order Processor ------------")

	orderItems := []string {
		"tshirt", "book", "hat", "mug_SALE",
	}

	var subtotal float64 = 0
	for _, item := range orderItems {
		price, found := calculateItemPrice(item)
		if found {
			subtotal += price
		}
	}

	fmt.Printf("Final price: %.2f\n", subtotal)
}
