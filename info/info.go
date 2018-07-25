package info

import "gopkg.in/yaml.v2"

type doc struct {
	Category    string
	Description string
	Message     string
}

// Map type for holding status code info
type codemap map[int]doc

// Info holds a map containing all c
type Info struct {
	fmt   Format
	codes codemap
}

// NewInfo returns an initalised *Info set to
// ROFF format
func NewInfo(y []byte) (*Info, error) {
	i := Info{}
	// Unmarshal and convert
	cm := make(codemap)
	err := yaml.Unmarshal(y, &cm)
	if err != nil {
		return nil, err
	}
	i.codes = cm
	i.fmt = ROFF
	return &i, nil
}

// Locate returns t/f if code c was found
func (m *Info) Locate(c int) bool {
	_, found := m.codes[c]
	return found
}

// SetFormat sets the output Format
// for a *Info
func (m *Info) SetFormat(f Format) {
	m.fmt = f
}
