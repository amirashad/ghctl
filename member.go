package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/google/go-github/v28/github"
)

func getMembers(org string, format string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	opt := &github.ListMembersOptions{ListOptions: github.ListOptions{PerPage: 100}}
	var objsAll []*github.User
	for {
		objs, resp, err := client.Organizations.ListMembers(ctx, org, opt)
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
		return *objsAll[i].Login < *objsAll[j].Login
	})

	if format == "normal" {
		for _, repo := range objsAll {
			fmt.Println(*repo.Login)
		}
	} else if format == "json" {
		bytes, _ := json.Marshal(objsAll)
		fmt.Println(string(bytes))
	}
}

func getPrimaryEmail() string {
	ctx := context.Background()
	client := createGithubClient(ctx)

	opt := &github.ListOptions{PerPage: 100}
	var objsAll []*github.UserEmail
	for {
		objs, resp, err := client.Users.ListEmails(ctx, opt)
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

	for _, e := range objsAll {
		if *e.Primary && *e.Verified {
			return *e.Email
		}
	}

	return ""
}
