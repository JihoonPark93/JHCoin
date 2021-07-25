package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/JihoonPark93/JHCoin/explorer"
	"github.com/JihoonPark93/JHCoin/rest"
)

func usage() {
	fmt.Printf("Welcome to JH conin\n\n")
	fmt.Printf("Please use the following command :\n\n")
	fmt.Printf("-port:	Set the PORT of the server\n")
	fmt.Printf("-mode:	Choose between 'html' and 'rest'\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	fmt.Println(*port, *mode)
	switch *mode {
	case "html":
		rest.Start(*port)
	case "rest":
		explorer.Start(*port)
	default:
		usage()
	}
}
