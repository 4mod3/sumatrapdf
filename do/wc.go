package main

import (
	"fmt"

	"github.com/kjk/u"
)

var srcFiles = u.MakeAllowedFileFilterForExts(".cpp", ".h", ".go")
var excludeDirs = u.MakeExcludeDirsFilter("ext")
var allFiles = u.MakeFilterAnd(excludeDirs, srcFiles)

func doLineCount() int {
	stats := u.NewLineStats()
	err := stats.CalcInDir("src", allFiles, true)
	if err != nil {
		fmt.Printf("doWordCount: stats.wcInDir() failed with '%s'\n", err)
		return 1
	}
	err = stats.CalcInDir("do", allFiles, true)
	if err != nil {
		fmt.Printf("doWordCount: stats.wcInDir() failed with '%s'\n", err)
		return 1
	}
	u.PrintLineStats(stats)
	return 0
}
