package info

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

type fmtPrint func(int) string

// Format is used as a type constant for
// the print methods
type Format int

// The Formats below define which print method
// a *Info will use
const (
	ROFF Format = iota
	SCDOC
	JSON
)

// PrintAll prints all codes in
// numerical order
func (m *Info) PrintAll() {
	m.PrintCodes(nil)
}

// PrintCodes takes an array of valid codes
// and prints them in numerical order
func (m *Info) PrintCodes(codes []int) {
	var s string
	switch m.fmt {
	case JSON:
		s = m.jsonPrint(codes)
	case SCDOC:
		s = m.fPrintf(m.scdPrint, codes)
	case ROFF:
		s = m.fPrintf(m.roffPrint, codes)
	}
	fmt.Println(s)
}

func (m *Info) fPrintf(fn fmtPrint, codes []int) string {
	var s string
	if codes == nil {
		codes = m.mapKeys()
	}
	sort.Ints(codes)
	for i := range codes {
		s += fn(codes[i])
	}
	return s
}

// prints information about code 'c'
// in man-page style roff markup.
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

// prints information about code 'c'
// in scdoc(1) markup.
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

// prints information about codes
// as a JSON representation
func (m *Info) jsonPrint(codes []int) string {
	var slice codemap
	if codes == nil {
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

// mapKeys returns the keys of m.codes
func (m *Info) mapKeys() []int {
	keys := make([]int, len(m.codes))
	i := 0
	for k := range m.codes {
		keys[i] = k
		i++
	}
	return keys
}
