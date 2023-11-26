package main

import (
	"flag"
	"fmt"
	"os"
)

const usage = `
`

func main() {
	var dsn string
	flag.StringVar(&dsn, "d", "database.sqlite", "database")
	flag.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}
}
