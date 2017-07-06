package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	s := "/Users/gyukebox/Golang/src/github.com/gyukebox/gopl/ch3/ch3.go"
	fmt.Println(basename(s))
	fmt.Println(comma("1234567890"))
	fmt.Println(intsToString([]int{1, 2, 3}))
	fmt.Println(isPalindrome("hello"))
	fmt.Println(isPalindrome("fzzf"))
}

func basename(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]

	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}

	return s
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(",")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func isPalindrome(s string) bool {
	n := len(s)
	if n <= 1 {
		return true
	}
	if s[0] != s[n-1] {
		return false
	}
	return isPalindrome(s[1 : n-2])
}
