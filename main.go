package main

import (
	"flag"
	"fmt"
	"os"
)

var orgFlag = flag.String("org", "", "Organisation name")
var versionFlag = flag.Bool("version", false, "App version")

const appVersion = "v0.0.2"

var token string
var org string

func main() {
	flag.Parse()

	if *versionFlag == true {
		fmt.Println(appVersion)
		return
	}

	args := flag.Args()
	// fmt.Println(args)
	// fmt.Println(*outputFlag)
	if len(args) < 2 {
		return
	}

	token = githubToken()
	org = githubOrg()

	if args[0] == "get" && args[1] == "repos" {
		outputFormat := getflag("-o", "normal", true)
		getRepos(org, outputFormat)
	}

	if args[0] == "get" && args[1] == "members" {
		outputFormat := getflag("-o", "normal", true)
		getMembers(org, outputFormat)
	}
}

func githubToken() string {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Println("Unauthorized: No token present. Please, add GITHUB_TOKEN environment variable")
		os.Exit(1)
	}
	return token
}

func githubOrg() string {
	if len(*orgFlag) == 0 {
		fmt.Println("org must be described!")
		os.Exit(1)
	}
	return *orgFlag
}
