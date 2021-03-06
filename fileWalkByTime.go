////////////////////////////////////////////////////////////////////////////
// Porgram: FileWalkByTime.go
// Purpose: Go file walk by time
// Authors: Tong Sun (c) 2017, All rights reserved
// Credits: berserkk
//          https://stackoverflow.com/a/44385800/2125837
////////////////////////////////////////////////////////////////////////////

package main

import (
	"os"
	"sort"

	"github.com/go-easygen/cli/clis"
)

type WalkFunc func(info os.FileInfo) error

type ByModTime []os.FileInfo

func (fis ByModTime) Len() int {
	return len(fis)
}

func (fis ByModTime) Swap(i, j int) {
	fis[i], fis[j] = fis[j], fis[i]
}

func (fis ByModTime) Less(i, j int) bool {
	return fis[i].ModTime().Before(fis[j].ModTime())
}

func fileWalkByTime(root string, walkFn WalkFunc) error {
	f, err := os.Open(root)
	clis.AbortOn("Open directory", err)
	fis, err := f.Readdir(-1)
	f.Close()
	clis.AbortOn("Read directory", err)
	// Get the actual file's FileInfo, instead of the symlink's
	// Readdir returns a slice of FileInfo values, as would be returned by Lstat
	// Lstat makes no attempt to follow the link, while Stat does
	for ii, fi := range fis {
		fis[ii], _ = os.Stat(fi.Name())
	}
	sort.Sort(ByModTime(fis))

	// https://godoc.org/os#FileInfo
	// https://godoc.org/time#pkg-constants
	for _, fi := range fis {
		// fmt.Println(fi.ModTime().Format(time.RFC3339), fi.Name())
		err = walkFn(fi)
	}
	return err
}
