package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	findThese := os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	interested := false
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "\t") {
			if interested {
				fmt.Println(txt)
			}
		} else {
			interested = false
			if !strings.Contains(txt, "STEXT") { // marks funcs
				continue
			}
			for _, this := range findThese {
				if strings.Contains(txt, this) {
					interested = true
					fmt.Println(txt)
					break
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

// might also want to filter out PCDATA by default, it's very noisy.
//  the go docs on what it is are... terse: https://golang.org/doc/asm
//  better described here: https://blog.altoros.com/golang-part-4-object-files-and-function-metadata.html
