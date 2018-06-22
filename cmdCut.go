////////////////////////////////////////////////////////////////////////////
// Program: picv
// Purpose: picture vault
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/go-easygen/cli"
	"github.com/go-easygen/cli/clis"
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

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	Opts OptsT
	cut  cuttingT
)

////////////////////////////////////////////////////////////////////////////
// cut

func cutCLI(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*cutT)
	imgRegex, _ = regexp.Compile(gcm[rootArgv.Case] + rootArgv.Glob + "$")
	Opts.DFN, Opts.Gap, Opts.Pod, Opts.Verbose =
		rootArgv.DFN, argv.Gap, argv.Pod, rootArgv.Verbose.Value()
	clis.Setup(progname, Opts.Verbose)
	ctx.JSON(Opts)
	fmt.Println()

	// visit the dirs provided on the command line
	return picVault(ctx.Args())
}

func picVault(dirs []string) error {

	//fmt.Println(dirs)
	for _, dir := range dirs {
		clis.Verbose(1, "Visting folder '%s'\n", dir)
		os.Chdir(dir)
		// open output file
		var err error
		cut.df, err = os.Create(Opts.DFN)
		clis.AbortOn("Creating directive file", err)
		err = fileWalkByTime(".", buildupPods)
		clis.AbortOn("File path walk", err)
		err = cut.df.Close()
		clis.AbortOn("Closing directive file", err)
	}

	return nil
}

func buildupPods(f os.FileInfo) error {
	cut.CutPods(f)
	return nil
}

func (cut *cuttingT) CutPods(f os.FileInfo) {
	// https://godoc.org/time
	// set lastDate to the first file date
	if !imgRegex.MatchString(f.Name()) {
		clis.Verbose(3, "File '%s' ignored", f.Name())
		return
	}
	fDay := f.ModTime().Format(dayFmt)
	if cut.picFiles == 0 {
		clis.Verbose(1, ">  %s, %s", f.Name(), fDay)
	}
	cut.picFiles++
	if cut.lastDate.IsZero() {
		cut.lastDate = f.ModTime()
		fmt.Fprintln(cut.df, fDay)
	}
	if cut.picFiles > Opts.Pod {
		correction := float32(cut.picFiles-Opts.Pod) / float32(Opts.Pod)
		correction *= correction
		if int(f.ModTime().Sub(cut.lastDate).Hours()) >
			int((float32(Opts.Gap)-correction)*24) {
			clis.Verbose(1, "<  %s, %s", f.Name(), fDay)
			// create a new pod
			fmt.Fprintln(cut.df, fDay)
			cut.picFiles = 0
		}
	}
	clis.Verbose(2, "%d: %s, %s", cut.picFiles, f.Name(), fDay)
	cut.lastDate = f.ModTime()
}
