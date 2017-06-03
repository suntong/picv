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

////////////////////////////////////////////////////////////////////////////
// cut

func cutCLI(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*cutT)
	Opts.Gap, Opts.Pod, Opts.Verbose = argv.Gap, argv.Pod, rootArgv.Verbose.Value()
	ctx.JSON(Opts)
	fmt.Println()

	return picVault()
}

func picVault() error {
	return nil
}
