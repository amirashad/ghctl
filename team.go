package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/google/go-github/v33/github"
)

func getTeams(org string, format string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	opt := &github.ListOptions{PerPage: 100}
	var objsAll []*github.Team
	for {
		objs, resp, err := client.Teams.ListTeams(ctx, org, opt)
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

func getTeam(org string, team *string, format string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	obj, _, err := client.Teams.GetTeamBySlug(ctx, org, *team)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if format == "normal" {
		fmt.Println(*obj.Name)
	} else if format == "json" {
		bytes, _ := json.Marshal(obj)
		fmt.Println(string(bytes))
	}
}

func addTeamToRepo(org string,
	repo, team, permission string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	perm := &github.TeamAddTeamRepoOptions{
		Permission: permission,
	}

	resp, err := client.Teams.AddTeamRepoBySlug(ctx, org, team, org, repo, perm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp.Status)
}
