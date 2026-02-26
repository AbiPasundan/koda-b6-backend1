package main

import (
	"satu/pkg"
)

var IsAuthenticated bool = false

func main() {
	if IsAuthenticated == true {
		pkg.Loggined()
	} else {
		pkg.Auth()
	}
}
