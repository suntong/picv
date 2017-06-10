////////////////////////////////////////////////////////////////////////////
// Program: picv
// Purpose: picture vault
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"github.com/mkideal/cli"
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
	Desc:   "picture vault\nbuilt on " + buildTime,
	Text:   "Tool to deal with camera pictures and put them in vault",
	Global: true,
	Argv:   func() interface{} { return new(rootT) },
	Fn:     picv,

	NumArg: cli.AtLeast(1),
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "picv"
//          VERSION   = "0.1.0"
//          buildTime = "2017-06-10"
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
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
// Main dispatcher

//  func picv(ctx *cli.Context) error {
//  	ctx.JSON(ctx.RootArgv())
//  	ctx.JSON(ctx.Argv())
//  	fmt.Println()

//  	return nil
//  }

// Template for CLI handling starts here

////////////////////////////////////////////////////////////////////////////
// cut

// func cutCLI(ctx *cli.Context) error {
// 	rootArgv = ctx.RootArgv().(*rootT)
// 	argv := ctx.Argv().(*cutT)
// 	fmt.Printf("[cut]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
// 	return nil
// }

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

// func archCLI(ctx *cli.Context) error {
// 	rootArgv = ctx.RootArgv().(*rootT)
// 	argv := ctx.Argv().(*archT)
// 	fmt.Printf("[arch]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
// 	return nil
// }

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
