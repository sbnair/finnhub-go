// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finnhub

import (
	"testing"
)

func TestExchange(t *testing.T) {

	tests := []struct {
		input Exchange
		want  string
	}{
		{input: ExchangeUS, want: "US"},
		{input: ExchangeUK, want: "L"},
	}

	for _, tc := range tests {
		got := tc.input.String()
		if tc.want != got {
			t.Fatalf("exchange not valid, want: %v got: %v", tc.want, got)
		}
	}
}
