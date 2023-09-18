package main

import (
	"fmt"
	"sync"

	"cryptomasters.com/go/api"
)

func main() {
	currencies := []string{"BTC", "ETH", "SOL", "DOT"}
	var wg sync.WaitGroup
	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			getCurrencyData(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()

}

func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err != nil {
		fmt.Printf("error on get rate")
	}
	fmt.Printf("The current price of %s is: %.2f\n", rate.Currency, rate.Price)
}
