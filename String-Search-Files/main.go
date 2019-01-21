package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

func main() {

	substr := flag.String("word", "of the", "String to Search")
	flag.Parse()

	var wg sync.WaitGroup
	workers := 10
	wg.Add(workers)

	ch := make(chan string)

	workingDir, _ := os.Getwd()
	filesDir := workingDir + "/" + "files/"

	files, err := ioutil.ReadDir(filesDir)
	if err != nil {
		fmt.Println("No Folder name 'files' present in Current Working Directory")
		os.Exit(1)
	}

	for w := 0; w < workers; w++ {
		go fileRead(*substr, ch, &wg)
	}

	cnt := 0
	for _, f := range files {
		ch <- filesDir + f.Name()
		cnt++
	}
	close(ch)
	wg.Wait()

	fmt.Println("\n-> Keyword Searched:'", *substr, "'")
	fmt.Println("-> Number of Files Searched: ", cnt, "files")
	fmt.Println("")
}

func fileRead(substr string, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		path, ok := <-ch
		if !ok {
			return
		}
		b, err := os.Open(path)
		if err != nil {
			fmt.Print(err)
		}
		defer b.Close()

		scanner := bufio.NewScanner(b)

		count := 0
		for scanner.Scan() {
			if strings.Contains(scanner.Text(), substr) {
				count++
			}
		}

		fmt.Println("\nThe file:", path, "has ", count, " occurances.")
	}
}
