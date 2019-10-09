package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v25/github"
	"golang.org/x/oauth2"
)

func createProtection(org, repoName, protectionPattern string, minApprove int, dismissStalePrApprovals, codeOwner bool) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// isGradle := isFileExists(ctx, client, org, repoName, branch, "build.gradle")
	// isMaven := isFileExists(ctx, client, org, repoName, branch, "pom.xml")
	// isGo := false   //isFileExists(ctx, client, org, repoName, branch, "go.mod")
	// isNode := false //isFileExists(ctx, client, org, repoName, branch, "package.json")

	// fmt.Println(repoName, "\t\tgradle:", isGradle, "maven:", isMaven, "go:", isGo, "node:", isNode)

	preq := &github.ProtectionRequest{
		RequiredPullRequestReviews: &github.PullRequestReviewsEnforcementRequest{
			RequireCodeOwnerReviews:      codeOwner,
			RequiredApprovingReviewCount: minApprove,
			DismissStaleReviews:          dismissStalePrApprovals,
		},
	}

	protection, _, err := client.Repositories.UpdateBranchProtection(ctx, org, repoName, protectionPattern, preq)
	if err == nil {
		fmt.Println(protection)
	} else {
		fmt.Println(err)
	}
}
