package fscli

import (
	"cmp"
	"fmt"
	"io/fs"
)

func (c *CLI) globMain(args []string) {
	if len(args) == 0 {
		fmt.Fprintf(c.stderr(), "usage: %s glob pattern [...]\n", cmp.Or(c.Name, "fscli"))
		c.exitCode = 1
		return
	}
	for _, arg := range args {
		c.glob(arg)
	}
}

func (c *CLI) glob(pattern string) {
	names, err := fs.Glob(c.FS, pattern)
	if err != nil {
		c.eprintln("glob", err)
		c.exitCode = 1
		return
	}
	for _, name := range names {
		fmt.Fprintln(c.stdout(), name)
	}
}
