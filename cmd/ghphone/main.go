// Command ghphone checks Ghanaian mobile numbers given as arguments or one per line on stdin.
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nadeem0305/ghphone/phone"
)

func main() {
	inputs := os.Args[1:]
	if len(inputs) == 0 {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			inputs = append(inputs, sc.Text())
		}
	}
	bad := 0
	for _, in := range inputs {
		n, err := phone.Parse(in)
		if err != nil {
			fmt.Printf("%-18s invalid\n", in)
			bad++
			continue
		}
		fmt.Printf("%-18s %s  %s\n", in, n.E164, n.Network)
	}
	if bad > 0 {
		os.Exit(1)
	}
}
