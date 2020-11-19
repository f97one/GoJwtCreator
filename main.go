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
	noLf := flag.Bool("nolf", false, "not append LF to end of line")
	help := flag.Bool("help", false, "Display usage and exit")

	flag.Parse()

	if *help {
		printUsage()
		os.Exit(0)
	}

	if len(*privateKeyFile) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "Private Key not specified.")
		printUsage()
		os.Exit(1)
	}

	if len(*issuer) == 0 {
		_, _ = fmt.Fprintln(os.Stderr, "issuer is required.")
		printUsage()
		os.Exit(1)
	}

	if *expires < 1 {
		_, _ = fmt.Fprintf(os.Stderr, "Too little expiration : %d\n", *expires)
		printUsage()
		os.Exit(1)
	}

	ret, err := createJWT(*issuer, *expires, *privateKeyFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	if *noLf {
		fmt.Print(ret)
	} else {
		fmt.Println(ret)
	}
	os.Exit(0)
}

func printUsage() {
	fmt.Printf("usage : %s {-expires [expires]} -issuer [issuer] {-nolf} -private [private]\n", os.Args[0])
	flag.PrintDefaults()
}
