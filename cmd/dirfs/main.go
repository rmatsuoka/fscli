package main

import (
	"flag"
	"os"

	"github.com/rmatsuoka/fscli"
)

var (
	root = flag.String("root", ".", "root dir")
)

func main() {
	flag.Parse()
	fscli.Main("dirfs", os.DirFS(*root), flag.Args())
}
