package fscli

import (
	"cmp"
	"fmt"
	"io"
	"strings"
)

func (c *CLI) prefix(command string) string {
	return cmp.Or(c.Name, "fscli") + ": " + command + ": "
}

func (c *CLI) eprintln(command string, a ...any) {
	b := new(strings.Builder)
	b.WriteString(c.Name)
	b.WriteString(": ")
	if command != "" {
		b.WriteString(command)
		b.WriteString(": ")
	}
	fmt.Fprintln(b, a...)
	io.WriteString(c.stderr(), b.String())
}

func (c *CLI) eprintf(command string, format string, a ...any) {
	b := new(strings.Builder)
	b.WriteString(c.Name)
	b.WriteString(": ")
	if command != "" {
		b.WriteString(command)
		b.WriteString(": ")
	}
	fmt.Fprintf(b, format, a...)
	io.WriteString(c.stderr(), b.String())
}
