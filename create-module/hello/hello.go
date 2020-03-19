package hello

import "rsc.io/quote"

//import "rsc.io/quote/v2" for  specific version

func Hello() string {
	return quote.Hello()
}
