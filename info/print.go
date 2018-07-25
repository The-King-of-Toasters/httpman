package info

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type fmtPrint func(int) string

type Format int

const (
	Roff Format = iota
	Scdoc
	Json
)

// prints all codes in
// numerical order
func (m *Info) PrintAll() {
	keys := make([]int, len(m.codes))
	i := 0
	for k := range m.codes {
		keys[i] = k
		i++
	}
	m.PrintCodes(keys)
}

// takes an array of valid codes
// and prints them in numerical order
func (m *Info) PrintCodes(codes []int) {
	var s string
	if m.fmt == Json {
		s = m.jsonPrint(codes)
	} else {
		sort.Ints(codes)
		if m.fmt == Scdoc {
			s = m.fPrintf(m.scdPrint, codes)
		} else {
			s = m.fPrintf(m.roffPrint, codes)
		}
	}
	fmt.Println(s)
}

func (m *Info) fPrintf(fn fmtPrint, codes []int) string {
	var s string
	for i := range codes {
		s += fn(codes[i])
	}
	return s
}

// prints information about code 'c'
// man-page style roff markup.
func (m *Info) roffPrint(c int) string {
	s := fmt.Sprintf(`.TH HTTPMAN 7
.SH STATUS CODE
.B %d
- %s
.SH CATEGORY
%s
.SH DESCRIPTION
%s
`,
		c,
		m.codes[c].Message,
		m.codes[c].Category,
		m.codes[c].Description)
	// Escape all dashes
	s = strings.Replace(s, "-", "\\-", -1)
	return s
}

func (m *Info) scdPrint(c int) string {
	return fmt.Sprintf(`HTTPMAN(7)
# STATUS CODE
*%d* - %s
# CATEGORY
%s
# DESCRIPTION
%s
`,
		c,
		m.codes[c].Message,
		m.codes[c].Category,
		m.codes[c].Description)
}

func (m *Info) jsonPrint(codes []int) string {
	var slice codemap
	if len(codes) == len(m.codes) {
		slice = m.codes
	} else {
		slice = make(codemap, len(codes))
		for i := range codes {
			slice[codes[i]] = m.codes[codes[i]]
		}
	}
	s, _ := json.Marshal(slice)
	return string(s)
}
