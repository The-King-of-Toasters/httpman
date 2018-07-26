// httpman - 2018 Stephen Gregoratto
package main

import (
	"fmt"
	"os"
	"strconv"

	"git.sgregoratto.me/The-King-of-Toasters/httpman/info"
	"github.com/pborman/getopt/v2"
)

// Getopt flags
var all, help, scd, json bool

func init() {
	getopt.Flag(&all, 'a', "Print info on every code")
	getopt.Flag(&help, 'h', "Print this help")
	// Formats
	getopt.Flag(&scd, 's', "Print in scdoc format")
	getopt.Flag(&json, 'j', "Print in json format")
	getopt.Parse()
}

func main() {
	args := getopt.Args()
	if (!all && len(args) == 0) || help {
		getopt.PrintUsage(os.Stderr)
		os.Exit(1)
	}

	cInfo := info.NewInfo()
	if scd {
		cInfo.SetFormat(info.SCDOC)
	} else if json {
		cInfo.SetFormat(info.JSON)
	}

	if all {
		cInfo.PrintAll()
	} else {
		var codes []int
		for i := range args {
			c, err := strconv.Atoi(args[i])
			if err != nil {
				error("code " + args[i] + " is not valid.")
			}

			if cInfo.Locate(c) {
				codes = append(codes, c)
			} else {
				error("Cannot find code " + args[i] + " in YAMLFILE.")
			}
		}
		cInfo.PrintCodes(codes)
	}
}

// Simple error handler to avoid importing "errors".
// Useful for avoiding piping problems to nroff
func error(e string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e)
}
