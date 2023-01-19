package main

import (
	"os"

	"github.com/atotto/clipboard"
)

func check(e error) {
	if e != nil {
		panic("make sure network drive is mapped")
	}
}

func main() {
	PFILE := "Z:/backup/documents/Personal/zzz.txt"
	dat, err := os.ReadFile(PFILE)
	check(err)
	clipboard.WriteAll(string(dat))
}
