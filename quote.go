// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finnhub

// Quote is a data structure containing the stock price of a symbol
// It's used by the Quote function.
type Quote struct {
	// Open price of the day
	Open float32 `json:"o"`
	// High price of the day
	High float32 `json:"h"`
	// Low price of the day
	Low float32 `json:"l"`
	// Current price
	Current float32 `json:"c"`
	// Previous close price
	Previous float32 `json:"pc"`
}
