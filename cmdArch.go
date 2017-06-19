////////////////////////////////////////////////////////////////////////////
// Program: picv
// Purpose: picture vault
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"bufio"
	"os"
	"regexp"
	"time"

	"github.com/mkideal/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const mode = 0711

type rootOptsT struct {
	DFN     string
	Verbose int
}

type archingT struct {
	scanner  *bufio.Scanner
	picFiles int
	lastDate string
	nextDate string
}

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	rootOpts rootOptsT
	arch     archingT
)

////////////////////////////////////////////////////////////////////////////
// arch

func archCLI(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	//argv := ctx.Argv().(*archT)
	// fmt.Printf("[arch]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())

	imgRegex, _ = regexp.Compile(gcm[rootArgv.Case] + rootArgv.Glob + "$")
	rootOpts.DFN, rootOpts.Verbose = rootArgv.DFN, rootArgv.Verbose.Value()
	// process the dirs provided on the command line
	return picArch(ctx.Args())
}

func picArch(dirs []string) error {
	for _, dir := range dirs {
		os.Chdir(dir)
		verbose(1, rootOpts.Verbose, "Visting folder '%s'\n", dir)
		// open file
		file, err := os.Open(rootOpts.DFN)
		abortOn("Opening directive file", err)
		// create a new scanner and read the file line by line
		arch.scanner = bufio.NewScanner(file)
		err = fileWalkByTime(".", createPods)
		abortOn("File path walk", err)
		err = file.Close()
		abortOn("Closing directive file", err)
	}
	return nil
}

func createPods(f os.FileInfo) error {
	return arch.ArchPods(f)
}

func (arch *archingT) ArchPods(f os.FileInfo) error {
	// https://godoc.org/time
	// set lastDate to the first file date
	if !imgRegex.MatchString(f.Name()) {
		verbose(3, rootOpts.Verbose, "File '%s' ignored", f.Name())
		return nil
	}
	// Get the actual file date, instead of the symlink's
	f, err := os.Stat(f.Name())
	abortOn("Get file stat", err)
	fName, fDay := f.Name(), f.ModTime().Format(dayFmt)
	if arch.lastDate == "" {
		arch.scanner.Scan()
		arch.lastDate = arch.scanner.Text()
		arch.enterNewPod()
		arch.scanner.Scan()
		arch.nextDate = arch.scanner.Text()
	}
	if fDay >= arch.nextDate {
		verbose(3, rootOpts.Verbose, "= %s, %s", arch.lastDate, arch.nextDate)
		// create in new pod
		arch.lastDate = arch.nextDate
		arch.enterNewPod()
		if arch.scanner.Scan() {
			arch.nextDate = arch.scanner.Text()
		} else {
			arch.nextDate = time.Now().Format(dayFmt)
		}
		verbose(3, rootOpts.Verbose, "= %s, %s", arch.lastDate, arch.nextDate)
	}
	verbose(2, rootOpts.Verbose, "%s: %s", fName, fDay)
	os.Symlink("../"+fName, arch.lastDate+"/"+fName)
	return nil
}

func (arch *archingT) enterNewPod() {
	os.Mkdir(arch.lastDate, mode)
	verbose(1, rootOpts.Verbose, "> %s", arch.lastDate)
}
