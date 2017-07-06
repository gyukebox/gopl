package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() //rune, nbytes, error 반환
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalud UTF-8 characters\n", invalid)
	}
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func removeDuplicate(strings []string) []string {
	for i := 0; i < len(strings)-1; i++ {
		if strings[i] == strings[i+1] {
			copy(strings[i:], strings[i+1:])
			strings = strings[:len(strings)-1]
		}
	}
	return strings
}

func rotate(a []int) {
	tmp := a[len(a)-1]
	for i := len(a) - 2; i >= 0; i-- {
		a[i+1] = a[i]
	}
	a[0] = tmp
}

func remove(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])
	return slice[:len(slice)-1]
}
