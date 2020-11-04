package main

import (
	"fmt"
	"github.com/piquette/finance-go/options"
	"github.com/piquette/finance-go/quote"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

func main(){
	iter := options.GetStraddle("TWTR")

	fmt.Println(iter.Meta())

	for iter.Next(){
		fmt.Println(iter.Straddle().Strike)
	}
	if iter.Err() != nil {
		fmt.Println(iter.Err())
	}

	// Basic quoute example.
	fmt.Println("--------------------------")
	fmt.Println("SPY quote data")
	q, err := quote.Get("SPY")
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(q)
	}

	// Basic chart example
	fmt.Println("----------------------------")
	fmt.Println("Basic chart example")
	params := &chart.Params{
		Symbol: "TWTR",
		Interval: datetime.OneHour,
	}
	iterr := chart.Get(params)

	for iterr.Next(){
		b := iterr.Bar()
		fmt.Println(datetime.FromUnix(b.Timestamp))
	}

	if iterr.Err() != nil {
		fmt.Println(iterr.Err())
	}

}