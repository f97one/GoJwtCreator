package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	privateKeyFile := flag.String("private", "", "Private Key file to sign JWT (required)")
	expires := flag.Int("expires", 30, "Time to expiration (in minutes, 30 minutes if not specified)")

	flag.Parse()

	if len(*privateKeyFile) == 0 {
		fmt.Println("Private Key not specified.")
		printUsage()
		os.Exit(1)
	}

	if *expires < 1 {
		fmt.Printf("Too little expiration : %d\n", *expires)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Printf("usage : %s -expires [expires] -private [private]\n", os.Args[0])
	flag.PrintDefaults()
}
