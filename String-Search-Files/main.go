package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {

	substr := flag.String("word", "of", "String to Search")
	flag.Parse()

	filechannel := make(chan string, 100)
	resultchannel := make(chan string, 100)

	workingDir, _ := os.Getwd()
	filesDir := workingDir + "/" + "files/"

	files, err := ioutil.ReadDir(filesDir)
	if err != nil {
		fmt.Println("No Folder name 'files' present in Current Working Directory")
		os.Exit(1)
	}

	for w := 1; w <= 10; w++ {
		go fileRead(filechannel, *substr, resultchannel)
	}

	path := ""

	for _, f := range files {
		path = filesDir + f.Name()
		filechannel <- path
	}

	close(filechannel)

	cnt := 0
	for range files {
		cnt++
		fmt.Print("#", cnt, " ", <-resultchannel)
	}
}

func fileRead(filechannel <-chan string, substr string, resultchannel chan<- string) {
	for path := range filechannel {
		b, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Print(err)
		}

		str := string(b)
		count := 0
		wordlist := strings.Fields(str)

		for _, word := range wordlist {
			if strings.Contains(word, substr) {
				count++
			}
		}

		result := "The file " + path + " word '" + substr + "' has " + strconv.Itoa(count) + " occurances\n\n"
		resultchannel <- result

	}
}
