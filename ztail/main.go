package main

import (
	"fmt"
	"os"
)

// parseInt converts a string to an integer manually.
func parseInt(s string) (int, error) {
	if len(s) == 0 {
		return 0, fmt.Errorf("empty string cannot be converted to int")
	}

	var num int
	sign := 1
	start := 0

	if s[0] == '-' {
		sign = -1
		start = 1
	} else if s[0] == '+' {
		start = 1
	}

	for i := start; i < len(s); i++ {
		ch := s[i]
		if ch < '0' || ch > '9' {
			return 0, fmt.Errorf("invalid character '%c' in number", ch)
		}
		num = num*10 + int(ch-'0')
	}

	return num * sign, nil
}

// numberOfBytes processes command-line arguments to extract the number of bytes and file names.
func numberOfBytes(args []string) (int, []string) {
	n := len(args)
	nbytes := 0
	var files []string

	for i, v := range args {
		if v == "-c" {
			if i >= n-1 {
				fmt.Println("tail: option requires an argument -- 'c'")
				fmt.Println("Try 'tail --help' for more information.")
				os.Exit(1)
			}
			arg := args[i+1]

			var err error
			nbytes, err = parseInt(arg)
			if err != nil {
				fmt.Printf("tail: invalid number of bytes: %s\n", arg)
				os.Exit(1)
			}
			continue
		}

		if _, err := parseInt(v); err != nil {
			files = append(files, v)
		}
	}

	return nbytes, files
}

// fileSize retrieves the size of the file.
func fileSize(fi *os.File) int64 {
	fileInfo, err := fi.Stat()
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return fileInfo.Size()
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: <program> -c <number_of_bytes> <file1> [<file2> ...]")
		os.Exit(1)
	}

	nbytes, files := numberOfBytes(os.Args[1:])

	if nbytes <= 0 {
		fmt.Println("tail: invalid number of bytes")
		os.Exit(1)
	}

	printName := len(files) > 1
	exitStatus := 0

	for j, fileName := range files {
		fi, err := os.Open(fileName)
		if err != nil {
			fmt.Printf("open %s: %s\n", fileName, err.Error())
			exitStatus = 1
			continue
		}
		defer fi.Close()

		if printName {
			if j > 0 {
				fmt.Println()
			}
			fmt.Printf("==> %s <==\n", fileName)
		}

		fileSize := fileSize(fi)
		if nbytes > int(fileSize) {
			nbytes = int(fileSize)
		}

		read := make([]byte, nbytes)
		_, err = fi.ReadAt(read, fileSize-int64(nbytes))
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Print(string(read))

		if j < len(files)-1 {
			fmt.Println()
		}
	}

	os.Exit(exitStatus)
}
