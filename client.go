package main

import (
	"context"

	"github.com/google/go-github/v82/github"
	"golang.org/x/oauth2"
)

func createGithubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: args.Token})
	tc := oauth2.NewClient(ctx, ts)

	var client *github.Client
	var err error

	if args.Host != "" {
		client, err = github.NewClient(tc).WithEnterpriseURLs(args.Host, args.Host)
		CheckIfError(err)
	} else {
		client = github.NewClient(tc)
	}

	return client
}
