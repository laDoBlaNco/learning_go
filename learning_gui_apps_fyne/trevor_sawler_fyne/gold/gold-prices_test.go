package main

import (
	"testing"
)

// our actual test
func TestGold_GetPrices(t *testing.T) {
	g := Gold{
		Prices: nil,
		Client: client,
	}

	p, err := g.GetPrices()
	if err != nil {
		t.Error(err)
	}

	// since its now using our custom client and not actually going to the network, we
	// can make a more useful test as we know exactly what json we should be getting back
	if p.Price != 2000.01 {
		t.Error("wrong price returned:", p.Price)
	}
}
