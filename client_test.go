// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finnhub

import (
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
			t.Fatalf("expected: %v got and error: %v", tc.want, err)
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

// func TestQuote(t *testing.T) {
// 	tests := []struct {
// 		input string
// 		want  string
// 	}{
// 		{input: "TSLA", want: "Tesla, Inc."},
// 		{input: "GOOG", want: "Alphabet Inc."},
// 	}

// 	client := NewClient("")
// 	for _, tc := range tests {
// 		quote, _ := client.Quote(tc.input)
// 		got := quote.Price
// 		if !reflect.DeepEqual(tc.want, got) {
// 			t.Fatalf("expected: %v, got: %v", tc.want, got)
// 		}
// 	}
// }
