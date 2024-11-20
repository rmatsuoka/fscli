package fscli

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
)

type walkCLI struct {
	*CLI
	lflag bool
	tflag bool
}

func (c *CLI) walkMain(args []string) {
	flagSet := flag.NewFlagSet("walk", flag.ContinueOnError)
	flagSet.SetOutput(c.stderr())

	walkc := &walkCLI{CLI: c}
	flagSet.BoolVar(&walkc.lflag, "l", false, "long format")
	flagSet.BoolVar(&walkc.tflag, "t", false, "print type")
	err := flagSet.Parse(args)
	if errors.Is(err, flag.ErrHelp) {
		c.exitCode = 1
		return
	}
	if err != nil {
		c.eprintln("walk", err)
		c.exitCode = 1
		return
	}
	if flagSet.NArg() == 0 {
		walkc.walk(".")
	} else {
		for _, arg := range flagSet.Args() {
			walkc.walk(arg)
		}
	}
}

func (c *walkCLI) walk(name string) {
	err := fs.WalkDir(c.FS, name, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			c.eprintln("walk", err)
			return nil
		}
		c.printFile(path, d)
		return nil
	})
	if err != nil {
		c.exitCode = 1
		c.eprintln("walk", err)
	}
}

func (c *walkCLI) printFile(path string, e fs.DirEntry) {
	if c.lflag {
		info, err := e.Info()
		if err != nil {
			c.eprintln("ls", err)
			c.exitCode = 1
		}
		fmt.Fprintln(c.stdout(), formatFileInfo(info, path))
		return
	}
	if c.tflag {
		// do like fs.FormatDirEntry(e) but print path instead of e.Name().
		buf := make([]byte, 0, len(path)+5)
		mode := e.Type().String()
		buf = append(buf, mode[:len(mode)-9]...)
		buf = append(buf, ' ')
		buf = append(buf, path...)
		if e.IsDir() {
			buf = append(buf, '/')
		}
		fmt.Fprintf(c.stdout(), "%s\n", buf)
		return
	}
	fmt.Fprintln(c.stdout(), path)
}
