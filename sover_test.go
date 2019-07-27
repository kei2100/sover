package main

import (
	"bytes"
	"strings"
	"testing"
	"unicode"
)

func TestSover(t *testing.T) {
	in := `
v1.0.0
v1.0.0-rc.1
v1.0.0-rc.0
v1.0.0-rc.1+200
v1.0.0-rc.1+100
v1.0.0-alpha
v1.2.0
v1.10.0
2.0.0
2.0.0-rc.0
`

	want := `
v1.0.0-alpha
v1.0.0-rc.0
v1.0.0-rc.1
v1.0.0-rc.1+200
v1.0.0-rc.1+100
v1.0.0
v1.2.0
v1.10.0
2.0.0-rc.0
2.0.0
`
	want = strings.TrimLeftFunc(want, unicode.IsSpace)

	stdin = bytes.NewBufferString(in)
	var got bytes.Buffer
	stdout = &got

	main()

	if g, w := got.String(), want; g != w {
		t.Errorf("\n--- got --\n%v\n--- want ---\n%v", g, w)
	}
}
