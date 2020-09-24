// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finnhub

// StockSymbol is a data structure containing the name of a company along with
// its stock ticker symbol and currency. It's used by the StockSymbols function
// to return a list of symbols for a given exchange.
type StockSymbol struct {
	Description   string `json:"description"`
	DisplaySymbol string `json:"displaySymbol"`
	Symbol        string `json:"symbol"`
	Type          string `json:"type"`
	Currency      string `json:"currency"`
}
