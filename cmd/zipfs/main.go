package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rmatsuoka/fscli"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("zipfs: ")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "usage: zipfs zipfile command ...")
		os.Exit(1)
	}

	r, err := zip.OpenReader(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	fscli.Main("zipfs", r, flag.Args()[1:])
}
