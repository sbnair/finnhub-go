// Copyright 2020 Tony Mackay.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package finnhub

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
)

const baseURL = "https://finnhub.io/api/v1"

// ErrLimitExceeded is returned when the API request limit has been reached.
var ErrLimitExceeded = errors.New("finnhub: limit is exceeded")

// A Client is an Finnhub client.
type Client struct {
	token string
}

// NewClient returns a Client that has methods to call the Finnhub API using
// the specified API token of the account.
//
// When no API token is specified, the Finnhub API can still be called but the amount
// of API calls is limited.
func NewClient(token string) Client {
	return Client{
		token: token,
	}
}

// StockSymbols returns stock symbols listed on the specified exchange.
func (c *Client) StockSymbols(exchange string) ([]StockSymbol, error) {
	result := []StockSymbol{}

	// build URL
	endpoint, err := url.Parse(baseURL)
	if err != nil {
		return result, err
	}

	endpoint.Path += "/stock/symbol"

	// Query params
	params := url.Values{}
	params.Add("exchange", exchange)
	params.Add("token", c.token)
	endpoint.RawQuery = params.Encode()

	resp, err := http.Get(endpoint.String())
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return result, ErrLimitExceeded
	}

	decodeErr := json.NewDecoder(resp.Body).Decode(&result)
	if decodeErr != nil {
		return result, decodeErr
	}

	return result, nil
}

// Quote returns the real-time quote data for US stocks.
func (c *Client) Quote(symbol string) (Quote, error) {
	result := Quote{}

	// build URL
	endpoint, err := url.Parse(baseURL)
	if err != nil {
		return result, err
	}

	endpoint.Path += "/quote"

	// Query params
	params := url.Values{}
	params.Add("symbol", symbol)
	params.Add("token", c.token)
	endpoint.RawQuery = params.Encode()

	resp, err := http.Get(endpoint.String())
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 429 {
		return result, ErrLimitExceeded
	}

	decodeErr := json.NewDecoder(resp.Body).Decode(&result)
	if decodeErr != nil {
		return result, decodeErr
	}

	return result, nil
}
