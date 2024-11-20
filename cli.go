package fscli

import (
	"cmp"
	"fmt"
	"io"
	"io/fs"
	"os"
)

type CLI struct {
	Name   string
	FS     fs.FS
	Stdout io.Writer
	Stderr io.Writer

	exitCode int
}

func (c *CLI) Usage() {
	fmt.Fprintf(c.stderr(), `usage: %s comamnd args....

command list:
  ls   [-d] [-l] [name ...]
  read [name ...]
  walk [-l] [-t] [root ...]
  glob pattern [...]
`, cmp.Or(c.Name, "fscli"))
}

func New(name string, fsys fs.FS) *CLI {
	if fsys == nil {
		panic("fsys is nil")
	}
	return &CLI{
		Name: name,
		FS:   fsys,
	}
}

func Main(name string, fsys fs.FS, args []string) {
	c := New(name, fsys)
	os.Exit(c.Main(args))
}

func (c *CLI) Main(args []string) (exitCode int) {
	if len(args) < 1 {
		c.Usage()
		return 1
	}

	c.exitCode = 0

	switch args[0] {
	case "ls":
		c.lsMain(args[1:])
	case "read":
		c.readMain(args[1:])
	case "glob":
		c.globMain(args[1:])
	case "walk":
		c.walkMain(args[1:])
	case "help", "-h", "--help":
		c.Usage()
		return 1
	default:
		c.eprintf("", "unknown command %s\n", args[0])
		c.Usage()
		return 1
	}

	return c.exitCode
}

func (c *CLI) stdout() io.Writer {
	if c.Stdout == nil {
		return os.Stdout
	}
	return c.Stdout
}

func (c *CLI) stderr() io.Writer {
	if c.Stderr == nil {
		return os.Stderr
	}
	return c.Stderr
}
