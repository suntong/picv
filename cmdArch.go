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

////////////////////////////////////////////////////////////////////////////
// arch

func archCLI(ctx *cli.Context) error {
	rootArgv := ctx.RootArgv().(*rootT)
	argv := ctx.Argv().(*archT)
	fmt.Printf("[arch]:\n  %+v\n  %+v\n  %v\n", rootArgv, argv, ctx.Args())
	return nil
}