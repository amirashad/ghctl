package main

import (
	"flag"
	"fmt"
	"os"
)

var orgFlag = flag.String("org", "", "Organisation name")
var versionFlag = flag.Bool("version", false, "App version")

const appVersion = "v0.0.8"

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

	outputFormat := getflag("-o", "normal", false)
	if args[0] == "get" && args[1] == "repos" {
		getRepos(org, outputFormat)
	} else if args[0] == "get" && args[1] == "members" {
		getMembers(org, outputFormat)
	} else if args[0] == "get" && args[1] == "teams" {
		getTeams(org, outputFormat)
	} else if args[0] == "create" && args[1] == "repo" {
		repo := flagsToRepo()
		createRepo(org, repo, outputFormat)
	} else if args[0] == "create" && args[1] == "branch" {
		repo := getflag("-repo", "", true)
		newBranch := getflag("-b", "", true)
		createBranch(org, repo, newBranch, outputFormat)
	} else if args[0] == "add" && args[1] == "file" {
		repo := getflag("-repo", "", true)
		branch := getflag("-b", "", true)
		files := getflag("-f", "", true)
		gitName := getflag("-gitname", "", true)
		gitEmail := getflag("-gitemail", "", true)
		commitmessage := getflag("-m", "Change "+files, false)
		addFiles(org, repo, branch, files, commitmessage, gitName, gitEmail, outputFormat)
	} else if args[0] == "create" && args[1] == "protection" {
		repo := getflag("-repo", "", true)
		protectionPattern := getflag("-p", "", true)
		minApprove := getintflag("-min-approve", 1, false)
		dismissStalePrApprovals := getboolflag("-dismiss-stale-pr-approvals", false, false)
		codeOwner := getboolflag("-code-owner", false, false)
		createProtection(org, repo, protectionPattern, minApprove, dismissStalePrApprovals, codeOwner)
	} else {
		fmt.Println("Error: not implemented yet") // TODO: show help
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
