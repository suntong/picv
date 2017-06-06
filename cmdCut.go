////////////////////////////////////////////////////////////////////////////
// Program: picv
// Purpose: picture vault
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type OptsT struct {
	Gap     int
	Pod     int
	Verbose int
}

var Opts OptsT

////////////////////////////////////////////////////////////////////////////
// cut

func cutCLI(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*cutT)
	Opts.Gap, Opts.Pod, Opts.Verbose = argv.Gap, argv.Pod, rootArgv.Verbose.Value()
	ctx.JSON(Opts)
	fmt.Println()

	// visit the dirs provided on the command line
	return picVault(ctx.Args())
}

func picVault(dirs []string) error {
	//fmt.Println(dirs)
	for _, dir := range dirs {
		verbose(1, Opts.Verbose, "Visting folder '%s'\n", dir)
		os.Chdir(dir)
		err := filepath.Walk(".", visitPic)
		abortOn("File path walk", err)
	}

	return nil
}

func visitPic(path string, f os.FileInfo, err error) error {
	// https://godoc.org/path/filepath#Walk
	// https://godoc.org/os#FileInfo
	if f.IsDir() {
		verbose(2, Opts.Verbose, "Ignoring directory entry '%s'", path)
	} else if strings.Contains(path, ".git/") {
		verbose(3, Opts.Verbose, "Ignoring git entry '%s'", path)
	} else {
		fmt.Printf("  File: %s with %d bytes\n", path, f.Size())
	}
	return nil
}
