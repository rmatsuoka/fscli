package fscli

import (
	"cmp"
	"fmt"
	"io"
)

func (c *CLI) readMain(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(c.stderr(), "usage: %s read name [...]\n", cmp.Or(c.Name, "fscli"))
		c.exitCode = 1
		return
	}
	for _, arg := range args {
		c.read(arg)
	}
}

func (c *CLI) read(name string) {
	f, err := c.FS.Open(name)
	if err != nil {
		c.eprintln("read", err)
	}
	defer f.Close()
	_, err = io.Copy(c.stdout(), f)
	if err != nil {
		c.eprintln("read", err)
	}
}
