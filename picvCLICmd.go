////////////////////////////////////////////////////////////////////////////
// Program: picv
// Purpose: picture vault
// Authors: Tong Sun (c) 2017, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"

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

//==========================================================================
// Main dispatcher

func picv(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	fmt.Println()
	ctx.JSON(ctx.Argv())
	fmt.Println()
	argv := ctx.Argv().(*rootT)
	Opts.Gap, Opts.Pod, Opts.Verbose = argv.Gap, argv.Pod, argv.Verbose.Value()
	ctx.JSON(Opts)
	fmt.Println()

	return nil
}
