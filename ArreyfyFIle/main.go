package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func RetrieveFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)

	if err != nil {
		return nil, err
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		return nil, statsErr
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes, err
}

func main() {
	var arrayStr string = ""
	file, err := RetrieveFile("FILE.IN")
	if err != nil {
		println("error")
		return
	}
	arrayStr += "["
	var fileLenght = len(file)
	for i := 0; i < fileLenght; i++ {
		var buffer = hex.EncodeToString(file[i : i+1])
		buffer = strings.ToUpper(buffer)
		buffer = "0x" + buffer
		arrayStr += buffer
		arrayStr += ","
	}
	arrayStr = strings.TrimRight(arrayStr, ",")
	arrayStr += "]"
	os.Remove("FILE.OUT")
	f, err := os.Create("FILE.OUT")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(arrayStr)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "array written successfully")
}
