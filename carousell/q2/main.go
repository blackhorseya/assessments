package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func conv(x, y int) int {
	if x < y {
		x, y = y, x
	}
	for x % y != 0 {
		x, y = y, x % y
	}
	return y
}

func findSmallestDivisor(s string, t string) int32 {
	gcd := conv(len(s), len(t))
	for i := 0; i < gcd; i++ {
		if s[i] != t[i] {
			return -1
		}
	}
	for i := range s {
		if s[i] != s[i%gcd] {
			return -1
		}
	}
	for i := range t {
		if t[i] != t[i%gcd] {
			return -1
		}
	}
	return int32(len(s[:gcd]))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

	s := readLine(reader)

	t := readLine(reader)

	result := findSmallestDivisor(s, t)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
