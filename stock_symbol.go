// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finnhub

// StockSymbol is a data structure containing the name of a company along with
// its stock ticker symbol and currency. It's used by the StockSymbols function
// to return a list of symbols for a given exchange.
type StockSymbol struct {
	// Symbol description (company or fund name etc)
	Description string `json:"description,omitempty"`
	// Display symbol name.
	DisplaySymbol string `json:"displaySymbol,omitempty"`
	// Unique symbol used to identify this symbol used in /stock/candle endpoint.
	Symbol string `json:"symbol,omitempty"`
	// Security type.
	Type string `json:"type,omitempty"`
	// Price's currency. This might be different from the reporting currency of fundamental data.
	Currency string `json:"currency,omitempty"`
}
