/*
HELP
Формат запуска: test.exe .../.../filename.txt
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFileString(filename string) (s int) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := bufio.NewReader(f)
	_, err = r.ReadString('\n')

	for err == nil {
		_, err = r.ReadString('\n')
		s++
	}
	if err == io.EOF {
		s = s + 1 // Последняя строка не разделяется /n, поэтому её прибавляем
	}
	if err != io.EOF {
		fmt.Println(err) // Пишем текст ошибки
		s = -1           // Возвращаем индикатор ошибки
	}
	return s
}

func main() {

	arg := os.Args[1]

	s := ReadFileString(arg)
	fmt.Print(s)
}
