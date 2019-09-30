package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func getflag(flg string, def string, fail bool) string {
	if !strings.HasPrefix(flg, "-") {
		return def
	}

	args := flag.Args()
	for i, f := range args {
		if f == flg {
			if len(args) <= i+1 || strings.HasPrefix(args[i+1], "-") {
				return failif(fail, def, "please provide output type from: [normal, wide, json]")
			}

			return args[i+1]
		}
	}

	return def
}

func failif(fail bool, def string, err string) string {
	if fail {
		fmt.Println(err)
		os.Exit(2)
	}

	return def
}
