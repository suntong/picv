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
	abortOn("Open directory", err)
	fis, err := f.Readdir(-1)
	f.Close()
	abortOn("Read directory", err)
	sort.Sort(ByModTime(fis))

	// https://godoc.org/os#FileInfo
	// https://godoc.org/time#pkg-constants
	for _, fi := range fis {
		// fmt.Println(fi.ModTime().Format(time.RFC3339), fi.Name())
		err = walkFn(fi)
	}
	return err
}
