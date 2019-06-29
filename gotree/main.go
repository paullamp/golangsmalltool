package main

import (
	"flag"
	"os"
)

func main() {
	dir, _ := os.Getwd()
	var maxlevel int
	flag.IntVar(&maxlevel, "maxlevel", -1, "level want to show")
	flag.Parse()
	walkDirs(dir, 0, maxlevel)
}
