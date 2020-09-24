# finnhub-go
A Go module that let's you access financial data using the Finnhub Stock API.

## Using
Add the module to your project using the go get command:

```
go get github.com/tonymackay/finnhub-go
```

## Examples
Get a list of stock symbols of companies listed on the NYSE and NASDAQ.

```
// create a new client with the given API token.
// using an empty string will limit the amount of API calls per day.
client := finnhub.NewClient("")

symbols, err := client.StockSymbols("US")
if err != nil {
    fmt.Printf("error: %v\n", err)
}

for _, s := range symbols {
    fmt.Printf("Name: %s Symbol: %s\n", s.Description, s.Symbol)
}
```

Get a stock quote for Tesla (TSLA)

```
client := finnhub.NewClient("")

symbols, err := client.Quote("TSLA")
if err != nil {
    fmt.Printf("error: %v\n", err)
}

fmt.Printf("Price: %d\n", s.Price)
```

## License
[MIT License](LICENSE)