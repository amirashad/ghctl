package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

func getRepos(org string, format string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	opt := &github.RepositoryListByOrgOptions{ /*Type: "private", */ ListOptions: github.ListOptions{PerPage: 100}}
	var objsAll []*github.Repository
	for {
		objs, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		objsAll = append(objsAll, objs...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	sort.Slice(objsAll, func(i, j int) bool {
		return *objsAll[i].Name < *objsAll[j].Name
	})

	if format == "normal" {
		for _, repo := range objsAll {
			fmt.Println(*repo.Name)
		}
	} else if format == "wide" {
		for _, repo := range objsAll {
			fmt.Println(repo.String())
		}
	} else if format == "json" {
		bytes, _ := json.Marshal(objsAll)
		fmt.Println(string(bytes))
	}
}

func flagsToRepo() github.Repository {
	return github.Repository{
		Name:        github.String(getflag("-n", "", true)),
		Description: github.String(getflag("-d", "", false)),
		Homepage:    github.String(getflag("-h", "", false)),

		Private:     github.Bool(getboolflag("-private", true, false)),
		HasIssues:   github.Bool(getboolflag("-issues", true, false)),
		HasProjects: github.Bool(getboolflag("-projects", true, false)),
		HasWiki:     github.Bool(getboolflag("-wikis", true, false)),

		AutoInit:          github.Bool(getboolflag("-a", false, false)),
		GitignoreTemplate: github.String(getflag("-g", "", false)),
		AllowMergeCommit:  github.Bool(getboolflag("-mergecommit", true, false)),
		AllowSquashMerge:  github.Bool(getboolflag("-squash", true, false)),
		AllowRebaseMerge:  github.Bool(getboolflag("-rebase", true, false)),
	}
}

func createRepo(org string, repo github.Repository, format string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	objs, _, err := client.Repositories.Create(ctx, org, &repo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if format == "normal" {
		fmt.Println(*objs.Name)
	} else if format == "wide" {
		fmt.Println(objs.String())
	} else if format == "json" {
		bytes, _ := json.Marshal(objs)
		fmt.Println(string(bytes))
	}
}
