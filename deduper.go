package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	flag.Parse()
	args := flag.Args()
	initialDirectory := args[0]
	finfos, err := ioutil.ReadDir(initialDirectory)
	if err != nil {
		log.Fatal(err)
	}
	filenames := make(map[string]bool, len(finfos))
	for _, f := range finfos {
		filenames[f.Name()] = true
	}
	for i := 1; i < len(args); i++ {
		files, err := ioutil.ReadDir(args[i])
		if err != nil {
			log.Fatal(err)
		}
		for _, f := range files {
			if _, ok := filenames[f.Name()]; ok {
				fmt.Printf("found a match! %s matches\n", f.Name())
			} else {
				//fmt.Printf("no match for %s\n", f.Name())
			}
		}
	}
}
