package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v33/github"
	"golang.org/x/oauth2"
)

func createGithubClient(ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: args.Token})
	tc := oauth2.NewClient(ctx, ts)

	var client *github.Client
	var err error

	if args.Enterprise {
		baseUrl := fmt.Sprintf("%s/api/v3/", args.EnterpriseUrl)
		uploadUrl := fmt.Sprintf("%s/api/uploads/", args.EnterpriseUrl)

		client, err = github.NewEnterpriseClient(baseUrl, uploadUrl, tc)
		CheckIfError(err)
	} else {
		client = github.NewClient(tc)
	}

	return client
}
