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
	"time"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

type OptsT struct {
	DFN     string
	Gap     int
	Pod     int
	Verbose int
}

type cuttingT struct {
	df       *os.File
	picFiles int
	lastDate time.Time
}

var (
	Opts OptsT
	cut  cuttingT
)

////////////////////////////////////////////////////////////////////////////
// cut

func cutCLI(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*cutT)
	Opts.DFN, Opts.Gap, Opts.Pod, Opts.Verbose =
		rootArgv.DFN, argv.Gap, argv.Pod, rootArgv.Verbose.Value()
	// ctx.JSON(Opts)
	// fmt.Println()

	// visit the dirs provided on the command line
	return picVault(ctx.Args())
}

func picVault(dirs []string) error {

	//fmt.Println(dirs)
	for _, dir := range dirs {
		verbose(1, Opts.Verbose, "Visting folder '%s'\n", dir)
		os.Chdir(dir)
		// open output file
		var err error
		cut.df, err = os.Create(Opts.DFN)
		abortOn("Creating directive file", err)
		err = filepath.Walk(".", createPods)
		abortOn("File path walk", err)
		err = cut.df.Close()
		abortOn("Closing directive file", err)
	}

	return nil
}

func createPods(path string, f os.FileInfo, err error) error {
	// https://godoc.org/path/filepath#Walk
	// https://godoc.org/os#FileInfo
	if f.IsDir() {
		verbose(2, Opts.Verbose, "Ignoring directory entry '%s'", path)
	} else if strings.Contains(path, ".git/") {
		verbose(3, Opts.Verbose, "Ignoring git entry '%s'", path)
	} else {
		//fmt.Printf("  File: %s with %d bytes\n", path, f.Size())
		cut.CutPods(f)
	}
	return nil
}

func (cut *cuttingT) CutPods(f os.FileInfo) {
	// https://godoc.org/time
	// set lastDate to the first file date
	if !imgRegex.MatchString(f.Name()) {
		return
	}
	cut.picFiles++
	verbose(1, Opts.Verbose, "%d: %s, %s",
		cut.picFiles, f.Name(), f.ModTime().Format(dayFmt))
	if cut.lastDate.IsZero() {
		cut.lastDate = f.ModTime()
	}
	if cut.picFiles > Opts.Pod &&
		int(f.ModTime().Sub(cut.lastDate).Hours()) > Opts.Gap*24 {
		// create a new pod
		fmt.Fprintln(cut.df, f.ModTime().Format(dayFmt))
		cut.picFiles = 0
	}
	cut.lastDate = f.ModTime()

}
