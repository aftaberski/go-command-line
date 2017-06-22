package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/wsxiaoys/terminal/color"
)

func ls(directory string) {

	f, err := os.Open(directory)
	defer f.Close()
	if err != nil {
		fmt.Printf("go-ls: couldn't access directory\n")
		return
	}

	files, err := f.Readdirnames(-1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "go-ls: couldn't read files in directory: %s\n", directory)
		return
	}
	for i := 0; i < len(files); i++ {
		fi, err := os.Stat(files[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		switch mode := fi.Mode(); {
		case mode.IsDir():
			color.Printf("@c%s\n", files[i])
		case mode.IsRegular():
			color.Printf("@m%s\n", files[i])
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		ls(".")
	} else {
		ls(flag.Args()[0])
	}
}
