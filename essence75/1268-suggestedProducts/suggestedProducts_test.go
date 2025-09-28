package _268_suggestedProducts

import "testing"

func TestSuggestedProducts(t *testing.T) {
	products := []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}
	searchProd := "mouse"

	suggestedProducts(products, searchProd)
}
