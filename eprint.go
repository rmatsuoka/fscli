package fscli

import (
	"bytes"
	"cmp"
	"fmt"
)

func (c *CLI) prefix(command string) string {
	return cmp.Or(c.Name, "fscli") + ": " + command + ": "
}

func (c *CLI) eprintln(command string, a ...any) {
	b := new(bytes.Buffer)
	b.WriteString(c.Name)
	b.WriteString(": ")
	b.WriteString(command)
	b.WriteString(": ")
	fmt.Fprintln(b, a...)
	c.stderr().Write(b.Bytes())
}

func (c *CLI) eprintf(command string, format string, a ...any) {
	b := new(bytes.Buffer)
	b.WriteString(c.Name)
	b.WriteString(": ")
	b.WriteString(command)
	b.WriteString(": ")
	fmt.Fprintf(b, format, a...)
	c.stderr().Write(b.Bytes())
}
