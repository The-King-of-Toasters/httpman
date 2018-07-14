// httpman - 2018 Stephen Gregoratto
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pborman/getopt/v2"
	"gopkg.in/yaml.v2"
)

// YAMLFILE describes where the HTTP status code YAML file
// is located. Should be patched per user needs
const YAMLFILE string = "/etc/httpman_codes.yml"

// global map to hold all status code info
// when YAMLFILE is unmarshaled
var CMAP = make(map[int]map[string]string)

// getopt flags
var all, help bool

func init() {
	getopt.Flag(&help, 'h', "Print this help")
	getopt.Flag(&all, 'a', "Print man page for every code")

	f, err := ioutil.ReadFile(YAMLFILE)
	if err != nil {
		critError("Could not read status code file. " +
			"Check if YAMLFILE is correct.")
	}
	err = yaml.Unmarshal(f, &CMAP)
	if err != nil {
		critError("Could not parse status code file. " +
			"Check if YAMLFILE is correctly formatted")
	}
}

func main() {
	getopt.Parse()
	args := getopt.Args()
	if (!all && len(args) == 0) || help {
		getopt.PrintUsage(os.Stderr)
		os.Exit(1)
	}
	
	var codes []int
	if all {
		for k := range CMAP {
			codes = append(codes, k)
		}
	} else {
		for i := range args {
			c, err := strconv.Atoi(args[i])
			if err != nil {
				error("code " + args[i] + " is not valid.")
			}

			if _, ok := CMAP[c]; !ok {
				error("Cannot find code " + args[i] + " in YAMLFILE.")
			} else {
				codes = append(codes, c)
			}
		}
	}
	order(codes)
}

// sort an array of ints and Print each one
func order(k []int) {
	sort.Ints(k)
	for _, v := range k {
		Print(v)
	}
}

// Print takes a valid key in CMAP and prints its
// associated values in man-page style
// troff markup.
func Print(c int) {
	s := fmt.Sprintf(`.TH HTTPMAN 7
.SH STATUS CODE
.B %d
- %s
.SH CATEGORY
%s
.SH DESCRIPTION
%s`,
		c, CMAP[c]["message"],
		CMAP[c]["category"],
		CMAP[c]["description"])
	// Escape all dashes
	s = strings.Replace(s, "-", "\\-", -1)
	fmt.Println(s)
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
