////////////////////////////////////////////////////////////////////////////
// Program: picv
// Purpose: picture vault
// Authors: Tong Sun (c) 2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

	"github.com/go-easygen/cli"
	"github.com/go-easygen/cli/clis"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

//==========================================================================
// picv

type rootT struct {
	cli.Helper
	Glob    string      `cli:"G,glob" usage:"glob defines the image files matching pattern" dft:".*\\.jpg"`
	Case    bool        `cli:"C,case-sensitive" usage:"case-sensitive for glob files pattern matching"`
	DFN     string      `cli:"d,dfn" usage:"directive file name, which contains the cutting directives" dft:"zz-directive"`
	Verbose cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name:   "picv",
	Desc:   "picture vault\nVersion " + version + " built on " + date,
	Text:   "Tool to deal with camera pictures and put them in vault",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     picv,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

// The OptsT type defines all the configurable options from cli.
//  type OptsT struct {
//  	Glob	string
//  	Case	bool
//  	DFN	string
//  	Verbose	cli.Counter
//  	Verbose int
//  }

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "picv"
//          version   = "0.1.0"
//          date = "2018-06-21"

//  	rootArgv *rootT
//  	// Opts store all the configurable options
//  	Opts OptsT
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,
//  		cli.Tree(cutDef),
//  		cli.Tree(archDef)).Run(os.Args[1:]); err != nil {
//  		fmt.Fprintln(os.Stderr, err)
//  	}
//  	fmt.Println("")
//  }

// Template for main dispatcher starts here
//==========================================================================
// Dumb root handler

//  func picv(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// cut

//  func cutCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*cutT)
//  	clis.Verbose(2, "[cut]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	Opts.Glob, Opts.Case, Opts.DFN, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Glob, rootArgv.Case, rootArgv.DFN, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	clis.Setup(progname, Opts.Verbose)
//  	//return nil
//  	return DoCut()
//  }
//
//  func DoCut() error {
//  	fmt.Printf("%s v %s. Separate picture into groups\n", progname, version)
//  	fmt.Println("Copyright (C) 2018, Tong Sun")
//  	return nil
//  }

type cutT struct {
	Gap int `cli:"g,gap" usage:"Gap in days to be considered as different group/vault" dft:"5"`
	Pod int `cli:"p,pod" usage:"Minimum number of picture to have before splitting to a different group/vault" dft:"15"`
}

var cutDef = &cli.Command{
	Name: "cut",
	Desc: "Separate picture into groups",

	Argv: func() interface{} { return new(cutT) },
	Fn:   cutCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}

////////////////////////////////////////////////////////////////////////////
// arch

//  func archCLI(ctx *cli.Context) error {
//  	rootArgv = ctx.RootArgv().(*rootT)
//  	argv := ctx.Argv().(*archT)
//  	clis.Verbose(2, "[arch]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
//  	Opts.Glob, Opts.Case, Opts.DFN, Opts.Verbose, Opts.Verbose =
//  		rootArgv.Glob, rootArgv.Case, rootArgv.DFN, rootArgv.Verbose, rootArgv.Verbose.Value()
//  	clis.Setup(progname, Opts.Verbose)
//  	//return nil
//  	return DoArch()
//  }
//
//  func DoArch() error {
//  	fmt.Printf("%s v %s. Archive groups of picture into vaults\n", progname, version)
//  	fmt.Println("Copyright (C) 2018, Tong Sun")
//  	return nil
//  }

type archT struct {
}

var archDef = &cli.Command{
	Name: "arch",
	Desc: "Archive groups of picture into vaults",

	Argv: func() interface{} { return new(archT) },
	Fn:   archCLI,

	NumArg:      cli.AtLeast(1),
	CanSubRoute: true,
}
