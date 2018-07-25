// httpman - 2018 Stephen Gregoratto
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"git.sgregoratto.me/The-King-of-Toasters/httpman/info"
	"github.com/pborman/getopt/v2"
)

// YAMLFILE describes where the HTTP status code YAML file
// is located. Should be patched per user needs
const YAMLFILE string = "/etc/httpman_codes.yml"

var (
	cinfo                *info.Info // Global Info
	all, help, scd, json bool       // Getopt flags
)

func init() {
	getopt.Flag(&all, 'a', "Print cinfo on every code")
	getopt.Flag(&help, 'h', "Print this help")
	// Formats
	getopt.Flag(&scd, 's', "Print in scdoc format")
	getopt.Flag(&json, 'j', "Print in json format")
	getopt.Parse()

	f, err := ioutil.ReadFile(YAMLFILE)
	if err != nil {
		critError("Could not read status code file. " +
			"Check if YAMLFILE is correct.")
	}
	cinfo, err = info.NewInfo(f)
	if err != nil {
		critError("Could not parse status code file. " +
			"Check if YAMLFILE is correctly formatted")
	}
	if scd {
		cinfo.SetFormat(info.SCDOC)
	} else if json {
		cinfo.SetFormat(info.JSON)
	}
}

func main() {
	args := getopt.Args()
	if (!all && len(args) == 0) || help {
		getopt.PrintUsage(os.Stderr)
		os.Exit(1)
	}

	if all {
		cinfo.PrintAll()
	} else {
		var codes []int
		for i := range args {
			c, err := strconv.Atoi(args[i])
			if err != nil {
				error("code " + args[i] + " is not valid.")
			}

			if cinfo.Locate(c) {
				codes = append(codes, c)
			} else {
				error("Cannot find code " + args[i] + " in YAMLFILE.")
			}
		}
		cinfo.PrintCodes(codes)
	}
}

// Simple error handler to avoid importing "errors".
// Useful for avoiding piping problems to nroff
func error(e string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", e)
}

// Error and die
func critError(e string) {
	error(e)
	os.Exit(1)
}
