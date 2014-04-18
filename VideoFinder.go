package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var videoFile [100]string
var pathContainer [100]string
var i int = 0

func visit(path string, f os.FileInfo, err error) error {

	sub := strings.Split(path, "/")
	for file, _ := range sub {
		if strings.Contains(sub[file], ".mp4") || strings.Contains(sub[file], ".mkv") || strings.Contains(sub[file], ".avi") {
			fmt.Println("[*] FOUND VIDEO FILE : " + sub[file])
			videoFile[i] = sub[file]
			pathContainer[i] = path
			i++
		}
	}

	return nil
}

func main() {
	var toFile string
	Datafile := "find.txt"
	root := flag.String("f", "/", "Start from where ?")
	flag.Parse()
	filepath.Walk(*root, visit)
	for i, _ := range pathContainer {
		if pathContainer[i] != "" {
			toFile += strconv.Itoa(i+1) + "$" + pathContainer[i] + "#\n"
		}
	}
	ioutil.WriteFile(Datafile, []byte(toFile), 0x777)
	for j := 0; j < i; j++ {
		fmt.Print("\n[*] " + videoFile[j])
	}
	fmt.Printf("\n[*]START FROM: %s\n", *root)
}
