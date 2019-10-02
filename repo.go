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
	var reposAll []*github.Repository
	for {
		repos, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		reposAll = append(reposAll, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	sort.Slice(reposAll, func(i, j int) bool {
		return *reposAll[i].Name < *reposAll[j].Name
	})

	if format == "normal" {
		for _, repo := range reposAll {
			fmt.Println(*repo.Name)
		}
	} else if format == "wide" {
		for _, repo := range reposAll {
			fmt.Println(repo.String())
		}
	} else if format == "json" {
		bytes, _ := json.Marshal(reposAll)
		fmt.Println(string(bytes))
	}
}
