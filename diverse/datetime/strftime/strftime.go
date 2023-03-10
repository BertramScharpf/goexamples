package main

import (
	"fmt"
	"time"

	strftime "github.com/jehiah/go-strftime"
	"rsc.io/quote"
)

func Hello() string {
	return quote.Hello()
}

func main() {
	t := time.Unix(1340244776, 0)
	utc, _ := time.LoadLocation("UTC")
	t = t.In(utc)
	fmt.Println(strftime.Format("%Y-%m-%d %H:%M:%S", t))
	// Output:
	// 2012-06-21 02:12:56
	//
	fmt.Println(Hello())
}
