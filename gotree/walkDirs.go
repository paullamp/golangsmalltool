package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func walkDirs(dirname string, level, maxlevel int) {
	// 1.open dir and find if is a dir , is not , quit
	if level > maxlevel && maxlevel != -1 {
		return
	}
	dirhandler, err := os.Open(dirname)
	if err != nil {
		log.Fatal("Openerr:", err)
	}
	handlerinfo, staterr := dirhandler.Stat()
	if staterr != nil {
		log.Fatal("Stat err:", staterr)
	}
	if !handlerinfo.IsDir() {
		fmt.Println("|--", handlerinfo.Name())
		return
	} else {
		fiInfos, err := dirhandler.Readdir(-1)
		if err != nil {
			log.Fatal(err)
		}
		for _, v := range fiInfos {
			if v.IsDir() {
				fmt.Println(strings.Repeat("   ", level), "|--\033[0;32m", v.Name(), "\033[0m")
				walkDirs(dirname+"/"+v.Name(), level+1, maxlevel)
			} else {

				fmt.Println(strings.Repeat("   ", level), "|--", v.Name())
			}

		}
	}

	dirhandler.Close()

}
