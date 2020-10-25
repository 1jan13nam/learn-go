package hello

import "rsc.io/quote/v3"

//Hello function
func Hello() string {
	return quote.HelloV3()
}

//Proverb testing
func Proverb() string {
	return quote.Concurrency()
}
