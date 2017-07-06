package main

import (
	"fmt"
)

func basename(s string) string {
	ret := ""
	for i := 0; i < len(s); i++ {
		if s[i] == '/' {
			ret = s[i+1:]
		}
	}
	return ret
}

func main() {
	dir := "/Users/gyukebox/school/2-1/pp/teamproject/elevator/elevator.c"
	basedir := basename(dir)
	fmt.Println(basedir)
}
