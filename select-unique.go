//This pgm writes unique strings from list of repeated strings separated by newlines, in a file given in commandline
package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
  if len(os.Args) == 1 {
		fmt.Printf("Usage: %s [path to the file]\n", os.Args[0])
		return
	}
	inputFile := os.Args[1:][0]

	//- reading file
	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Error in reading file", err)
	}

	inputArr := strings.Split(string(b), "\n")

	//making memory
	var memory []string
	for _, line := range inputArr {
		if line != "" {
			if !stringInSlice(line, memory) {
				memory = append(memory, line)
			}
		}
	} //memory full

	//- writing blank to file
	err = ioutil.WriteFile(inputFile, []byte(""), 0666)
	if err != nil {
		log.Fatal("Error in writing empty to file: ", err)
	}

	//- writing memory to file
	file, err := os.OpenFile(inputFile, os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Error in writing memory to file : ", err)
	}
	defer file.Close()

	for _, line := range memory {
		file.WriteString(line + "\n")
	}
	//Done
}

//@ https://github.com/peeyushsrj/golang-snippets/blob/master/string-in-slice.go
func stringInSlice(a string, b []string) bool {
	for _, el := range b {
		if strings.Contains(a, el) {
			return true
		}
	}
	return false
}
