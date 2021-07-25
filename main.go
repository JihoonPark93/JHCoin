package main

import (
	"github.com/JihoonPark93/JHCoin/cli"
	"github.com/JihoonPark93/JHCoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
