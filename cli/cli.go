package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/JihoonPark93/JHCoin/explorer"
	"github.com/JihoonPark93/JHCoin/rest"
)

func usage() {
	fmt.Printf("Welcome to JH conin\n\n")
	fmt.Printf("Please use the following command :\n\n")
	fmt.Printf("-port:	Set the PORT of the server\n")
	fmt.Printf("-mode:	Choose between 'html' and 'rest'\n")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	fmt.Println(*port, *mode)
	switch *mode {
	case "html":
		explorer.Start(*port)
	case "rest":
		rest.Start(*port)
	default:
		usage()
	}

}
