/*
HELP
Enter file path like: test2.exe ../../file
*/

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	Path := os.Args[1]
	file, err := ioutil.ReadFile(Path)

	if len(os.Args) <= 1 || len(os.Args) > 2 {
		fmt.Println("Enter file path like: test2.exe ../../file")
		return
	}

	lines_num := strings.Split(string(file), "\n")
	fmt.Println(len(lines_num))

	if err != nil {
		fmt.Println("Invalid file path: ", err)
		return
	}
}
