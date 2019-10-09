package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getflag(flg string, def string, fail bool) string {
	if !strings.HasPrefix(flg, "-") {
		return def
	}

	args := flag.Args()
	found := false
	for i, f := range args {
		if f == flg {
			found = true
			if len(args) <= i+1 || strings.HasPrefix(args[i+1], "-") {
				return failif(fail, def, "please provide output type from: [normal, wide, json]")
			}

			return args[i+1]
		}
	}
	if !found {
		return failif(fail, def, "please provide name of repo with", flg, "flag")
	}
	return def
}

func getboolflag(flg string, def bool, fail bool) bool {
	flag := getflag(flg, "", fail)
	v, err := strconv.ParseBool(flag)
	if err == nil {
		return v
	}
	return def
}

func getintflag(flg string, def int, fail bool) int {
	flag := getflag(flg, "", fail)
	v, err := strconv.ParseInt(flag, 10, 0)
	if err == nil {
		return int(v)
	}
	return def
}

func failif(fail bool, def string, err ...string) string {
	if fail {
		fmt.Println(err)
		os.Exit(2)
	}

	return def
}
