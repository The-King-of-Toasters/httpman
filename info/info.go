package info

//import "gopkg.in/yaml.v2"

// Info holds a map containing all
type Info struct {
	fmt   Format
	codes codemap
}

// NewInfo returns an initialised *Info set to
// ROFF format
func NewInfo() *Info {
	i := Info{codes: cm, fmt: ROFF}
	return &i
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
