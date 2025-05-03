//go:build unit
// +build unit

package entities

import (
	"testing"
)

func TestBinanceSymbol(t *testing.T) {
	if BinanceSymbol("ETH-USDC") != "ETHUSDC" {
		t.Error("Not equal")
	}
}
