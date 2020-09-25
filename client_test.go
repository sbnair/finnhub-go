// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finnhub

import (
	"fmt"
	"testing"
)

func TestSymbols(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "US", want: "TSLA"},
		{input: "L", want: "VWRL.L"},
	}

	client := NewClient("")
	for _, tc := range tests {
		symbols, err := client.StockSymbols(tc.input)
		if err != nil {
			t.Fatalf("error: %v", err)
			break
		}

		var found = false
		for _, v := range symbols {
			if v.Symbol == tc.want {
				found = true
				break
			}
		}

		if !found {
			t.Fatalf("expected: %v in list of symbols, got nothing", tc.want)
		}
	}
}

func TestQuote(t *testing.T) {
	tests := []struct {
		input string
	}{
		{input: "TSLA"},
		{input: "VWRL.L"},
	}

	client := NewClient("")
	for _, tc := range tests {
		quote, err := client.Quote(tc.input)
		if err != nil {
			t.Fatalf("error: %v", err)
			break
		}

		fmt.Printf("Quote(%s) Result: %f\n", tc.input, quote.Current)
		if quote.Current == 0 {
			t.Fatalf("expected none zero price for symbol: %v", tc.input)
		}
	}
}
