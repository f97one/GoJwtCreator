package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	privateKeyFile := flag.String("private", "", "Private Key file to sign JWT (required)")
	issuer := flag.String("issuer", "", "issuer string (required)")
	expires := flag.Int("expires", 30, "Time to expiration (in minutes, 30 minutes if not specified)")

	flag.Parse()

	if len(*privateKeyFile) == 0 {
		fmt.Println("Private Key not specified.")
		printUsage()
		os.Exit(1)
	}

	if len(*issuer) == 0 {
		fmt.Println("issuer is required.")
		printUsage()
		os.Exit(1)
	}

	if *expires < 1 {
		fmt.Printf("Too little expiration : %d\n", *expires)
		printUsage()
		os.Exit(1)
	}

	ret, err := createJWT(*issuer, *expires, *privateKeyFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(ret)
	os.Exit(0)
}

func printUsage() {
	fmt.Printf("usage : %s {-expires [expires]} -issuer [issuer] -private [private]\n", os.Args[0])
	flag.PrintDefaults()
}
