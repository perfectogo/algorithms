package main

import "fmt"

func main() {
	var (
		fnum1, fnum2 = 0, 1
	    n int = 10
	    i, fnumN int
	)

	for i = 0; i <= n; i++ {
        fnumN = fnum1
        fnum1 = fnum2
        fnum2 = fnum1 + fnumN

        fmt.Println(i, "; ", fnumN)
	}

	fmt.Printf("\n%d - fibonacci number is %d", n, fnumN)
}