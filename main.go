package main

import (
	"os"

	"github.com/jessevdk/go-flags"
	"github.com/svenfinke/gls/lib/dirwalk"
	"github.com/svenfinke/gls/lib/types"
)

var options types.Options

var parser = flags.NewParser(&options, flags.Default)

func main() {

	if _, err := parser.Parse(); err != nil {
		switch flagsErr := err.(type) {
		case flags.ErrorType:
			if flagsErr == flags.ErrHelp {
				os.Exit(0)
			}
			os.Exit(1)
		default:
			os.Exit(1)
		}
	}

	if options.Args.File == "" {
		options.Args.File = "."
	}

	dirwalk.Walk(options)
}
