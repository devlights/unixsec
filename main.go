package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Println("no arguments\tusage:: unixsec unix-seconds")
		os.Exit(1)
	}

	v := args[0]
	unixSec, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		fmt.Printf("invalid value [%v]\n", v)
		os.Exit(1)
	}

	fmt.Printf("%v\t-->\t%v\n", unixSec, time.Unix(unixSec, 0))

	os.Exit(0)
}
