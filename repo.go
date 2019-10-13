package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

func getRepos(org string, format string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken()})
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
	} else if format == "json" {
		bytes, _ := json.Marshal(objsAll)
		fmt.Println(string(bytes))
	}
}

func createRepo(org string,
	name, descr, homepage *string,
	private, noIssues, noProjects, noWiki, autoinit *bool,
	gitIgnoreTemplate, licenseTemplate *string,
	noMergeCommit, noSquashMerge, noRebaseMerge *bool,
	format string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: githubToken()})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	repo := &github.Repository{
		Name:        name,
		Description: descr,
		Homepage:    homepage,

		Private:     private,
		HasIssues:   not(noIssues),
		HasProjects: not(noProjects),
		HasWiki:     not(noWiki),
		AutoInit:    autoinit,

		GitignoreTemplate: gitIgnoreTemplate,
		LicenseTemplate:   licenseTemplate,

		AllowMergeCommit: not(noMergeCommit),
		AllowSquashMerge: not(noSquashMerge),
		AllowRebaseMerge: not(noRebaseMerge),
	}

	objs, _, err := client.Repositories.Create(ctx, org, repo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if format == "normal" {
		fmt.Println(*objs.Name)
	} else if format == "json" {
		bytes, _ := json.Marshal(objs)
		fmt.Println(string(bytes))
	}
}

func not(o *bool) *bool {
	result := true
	if o == nil {
		return &result
	}
	result = !*o
	return &result
}
