package main

import (
	"github.com/JihoonPark93/JHCoin/explorer"
	"github.com/JihoonPark93/JHCoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
