package walker //Change Package Name For import

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var File []string
var Path []string
var typeFile *string

func main() {
	Alone()
}

func videoCrawl(path string, f os.FileInfo, err error) error {

	sub := strings.Split(path, "\\")
	for i := 0; i < len(sub); i++ {

		if strings.Contains(sub[i], ".mp4") || strings.Contains(sub[i], ".mkv") || strings.Contains(sub[i], ".avi") {

			if strings.Contains(sub[i], ".torrent") == false && strings.Contains(sub[i], ".lnk") == false {
				//fmt.Println("[*] FOUND VIDEO FILE : " + sub[i])
				File = append(File, sub[i])
				Path = append(Path, path) // COntains the same things then VideoFile slice
			}
		}
	}
	return nil
}

func textCrawl(path string, f os.FileInfo, err error) error {

	sub := strings.Split(path, "\\")
	for i := 0; i < len(sub); i++ {

		if strings.Contains(sub[i], ".txt") || strings.Contains(sub[i], ".doc") || strings.Contains(sub[i], ".nfo") {
			fmt.Println("[*] FOUND TEXT FILE : " + sub[i])
			File = append(File, sub[i])
			Path = append(Path, path) // COntains the same things then VideoFile slice
		}
	}
	return nil
}

func costCrawl(path string, f os.FileInfo, err error) error {

	sub := strings.Split(path, "\\")
	for i := 0; i < len(sub); i++ {

		if sub[i] == *typeFile {

			fmt.Println("[*] FOUND FILE : " + sub[i])
			File = append(File, sub[i])
			Path = append(Path, path) // COntains the same things then VideoFile slice

		}
	}
	return nil
}

func writeLog(Datafile *string) {
	var toFile string
	for i, _ := range File {
		if File[i] != "" {
			toFile += "FILE #" + strconv.Itoa(i+1) + " -> " + File[i] + "\n"
		}
	}
	toFile += "\n #################### PATH ##################### \n"

	for j, _ := range Path {
		if Path[j] != "" {
			toFile += "[*]PATH => " + Path[j] + "\n"
		}
	}
	ioutil.WriteFile(*Datafile, []byte(toFile), 0x777)
}

func flashFile(root *string) {
	for i := 0; i < len(File); i++ {
		fmt.Print("\n[*] " + File[i])
	}
	fmt.Printf("\n[*]START FROM: %s\n", *root)
}

func Alone() {
	root := flag.String("from", "C:\\", "Start from where ?")
	logName := flag.String("log", "log.txt", "Log file name ?")
	typeFile = flag.String("type", "video", "-type video or text")
	flag.Parse()

	if *typeFile == "video" {
		filepath.Walk(*root, videoCrawl)
		writeLog(logName)
		//flashFile(root)
	} else if *typeFile == "text" {
		filepath.Walk(*root, textCrawl)
		writeLog(logName)
		//flashFile(root)
	} else {
		filepath.Walk(*root, costCrawl)
		writeLog(logName)
	}
}

func Impor(kind string, place string) []string {
	if kind == "video" {
		filepath.Walk(place, videoCrawl)
	} else if kind == "text" {
		filepath.Walk(place, textCrawl)
	} else {
		filepath.Walk(kind, costCrawl)
	}
	return File
}
