////////////////////////////////////////////////////////////////////////////
// Program: picv
// Purpose: picture vault
// Authors: Tong Sun (c) 2017-2018, All rights reserved
////////////////////////////////////////////////////////////////////////////

package main

//go:generate sh -v picv_cliGen.sh

import (
	"fmt"
	"os"
	"regexp"

	"github.com/go-easygen/cli"
)

////////////////////////////////////////////////////////////////////////////
// Constant and data type/structure definitions

const dayFmt = "2006-01-02"

// glob case map
type gcmT map[bool]string

////////////////////////////////////////////////////////////////////////////
// Global variables definitions

var (
	progname = "picv"
	version  = "0.2.1"
	date     = "2018-06-21"

	imgRegex *regexp.Regexp
	gcm      gcmT
)

////////////////////////////////////////////////////////////////////////////
// Function definitions

// Function main
func main() {
	gcm = gcmT{false: "", true: `(?i)`}

	// cli.SetUsageStyle(cli.DenseNormalStyle) // left-right, for up-down, use ManualStyle
	cli.SetUsageStyle(cli.ManualStyle) // up-down style
	//NOTE: You can set any writer implements io.Writer
	// default writer is os.Stdout
	if err := cli.Root(root,
		cli.Tree(cutDef),
		cli.Tree(archDef)).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println("")
}

//==========================================================================
// Main dispatcher

func picv(ctx *cli.Context) error {
	ctx.JSON(ctx.RootArgv())
	ctx.JSON(ctx.Argv())
	fmt.Println()

	return nil
}
