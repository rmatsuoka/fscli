package fscli

import (
	"cmp"
	"errors"
	"flag"
	"fmt"
	"io/fs"
)

type lsCLI struct {
	*CLI
	lflag bool
	dflag bool
}

func (c *CLI) lsMain(args []string) {
	flagSet := flag.NewFlagSet("ls", flag.ContinueOnError)
	flagSet.SetOutput(c.Stderr)

	lsc := &lsCLI{CLI: c}
	flagSet.BoolVar(&lsc.lflag, "l", false, "long option")
	flagSet.BoolVar(&lsc.dflag, "d", false, "print dir itself instead of its content")

	err := flagSet.Parse(args)
	if errors.Is(err, flag.ErrHelp) {
		c.exitCode = 1
		return
	}
	if err != nil {
		c.eprintln("ls", err)
		c.exitCode = 1
		return
	}

	if flagSet.NArg() == 0 {
		lsc.ls(".")
	} else {
		for _, arg := range flagSet.Args() {
			lsc.ls(arg)
		}
	}
}

func (c *lsCLI) ls(fname string) {
	stat, err := fs.Stat(c.FS, fname)
	if err != nil {
		c.eprintln("ls", err)
		c.exitCode = 1
		return
	}
	if c.dflag || !stat.IsDir() {
		c.printFile(fs.FileInfoToDirEntry(stat))
		return
	}

	es, err := fs.ReadDir(c.FS, fname)
	if err != nil {
		c.eprintln("ls", err)
		c.exitCode = 1
		return
	}

	for _, e := range es {
		c.printFile(e)
	}
}

func (c *lsCLI) printFile(e fs.DirEntry) {
	if c.lflag {
		info, err := e.Info()
		if err != nil {
			c.eprintln("ls", err)
			return
		}
		fmt.Fprintln(c.stdout(), formatFileInfo(info, ""))
		return
	}
	fmt.Fprintln(c.stdout(), e.Name())
}

func formatFileInfo(info fs.FileInfo, path string) string {
	return fmt.Sprintf("%-12s %12d %s %s",
		info.Mode(),
		info.Size(),
		info.ModTime().Format("2006-01-02 15:04:05"),
		cmp.Or(path, info.Name()),
	)
}
