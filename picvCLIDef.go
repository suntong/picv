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
	Gap     int         `cli:"g,gap" usage:"Gap in days to be considered as different group/vault" dft:"5"`
	Pod     int         `cli:"p,pod" usage:"Minimum number of picture to have before splitting to a different group/vault" dft:"15"`
	Verbose cli.Counter `cli:"v,verbose" usage:"Verbose mode (Multiple -v options increase the verbosity.)"`
}

var root = &cli.Command{
	Name: "picv",
	Desc: "picture vault\nbuilt on " + buildTime,
	Text: "Tool to deal with camera pictures and put them in vault" +
		"\n\nUsage:\n  picv [Options] dir [dirs...]",
	Argv: func() interface{} { return new(rootT) },
	Fn:   picv,

	CanSubRoute: true,
}

// Template for main starts here
////////////////////////////////////////////////////////////////////////////
// Global variables definitions

//  var (
//          progname  = "picv"
//          VERSION   = "0.1.0"
//          buildTime = "2017-06-03"
//  )

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
//  func main() {
//  	cli.SetUsageStyle(cli.ManualStyle) // up-down, for left-right, use NormalStyle
//  	//NOTE: You can set any writer implements io.Writer
//  	// default writer is os.Stdout
//  	if err := cli.Root(root,).Run(os.Args[1:]); err != nil {
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
