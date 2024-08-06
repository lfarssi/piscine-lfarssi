package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: <program> <number_of_bytes> <file1> [<file2> ...]")
		return
	}

	numBytes, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	if len(args) < 2 {
		fmt.Println("Error: File name missing")
		return
	}

	for _, fileName := range args[1:] {
		ztail(fileName, numBytes)
	}
}

func ztail(fileName string, numBytes int) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error opening file '%s': %s\n", fileName, err.Error())
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Printf("Error getting file info for '%s': %s\n", fileName, err.Error())
		return
	}

	fileSize := fileInfo.Size()
	if numBytes > int(fileSize) {
		numBytes = int(fileSize)
	}

	_, err = file.Seek(-int64(numBytes), io.SeekEnd)
	if err != nil {
		fmt.Printf("Error seeking file '%s': %s\n", fileName, err.Error())
		return
	}
	data := make([]byte, numBytes)
	_, err = file.Read(data)
	if err != nil && err != io.EOF {
		fmt.Printf("Error reading file '%s': %s\n", fileName, err.Error())
		return
	}

	fmt.Printf("%s", data)
}
