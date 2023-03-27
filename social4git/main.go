package main

import (
	"fmt"
	"os"

	"github.com/gov4git/lib4git/must"
	"github.com/social4git/social4git/social4git/cmd"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(must.Error); ok {
				fmt.Fprintln(os.Stderr, e)
				fmt.Fprintln(os.Stderr, string(e.Stack))
			} else {
				fmt.Fprintln(os.Stderr, r)
			}
		}
	}()
	cmd.Execute()
}
