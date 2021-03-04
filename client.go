package main

import (
	"context"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

func createGithubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: args.Token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
