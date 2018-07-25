package info

import "gopkg.in/yaml.v2"

type doc struct {
	Category    string
	Description string
	Message     string
}

// Map type for holding status code info
type codemap map[int]doc

type Info struct {
	fmt   Format
	codes codemap
}

// NewInfo returns an initialised *Info
func NewInfo(y []byte) (*Info, error) {
	i := Info{}
	// Unmarshal and convert
	cm := make(codemap)
	err := yaml.Unmarshal(y, &cm)
	if err != nil {
		return nil, err
	}
	i.codes = cm
	i.fmt = Roff
	return &i, nil
}

// Returns t/f if code c was found in Doc
func (m *Info) Locate(c int) bool {
	_, found := m.codes[c]
	return found
}

func (m *Info) SetFormat(f Format) {
	m.fmt = f
}
