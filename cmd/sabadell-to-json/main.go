package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/c032/go-sabadell"
)

func main() {
	flagset := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	flagset.Usage = func() {
		fmt.Fprintf(flagset.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flagset.Output(), "\n")
		fmt.Fprintf(flagset.Output(), "    %s FILE\n", os.Args[0])

		flagset.PrintDefaults()
	}
	flagset.Parse(os.Args[1:])

	args := flagset.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Expected to receive one file.\n")

		os.Exit(2)
	} else if len(args) > 1 {
		fmt.Fprintf(os.Stderr, "Only one file was expected, but received %d files.\n", len(args))

		os.Exit(2)
	}

	file := args[0]
	enc := json.NewEncoder(os.Stdout)

	lines, err := sabadell.ParseTXTFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())

		os.Exit(1)
	}

	for _, line := range lines {
		err := enc.Encode(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
	}
}
