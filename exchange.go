// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finnhub

// Exchange is used by the StockSymbols function
type Exchange string

const (
	// ExchangeUS to get symbols of companies listed on the NYSE and Nasdaq
	ExchangeUS = "US"
	// ExchangeUK is used to get symbols of companies listed on the LSE
	ExchangeUK = "L"
)

func (e Exchange) String() string {
	return string(e)
}
