package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
	"golang.org/x/oauth2"
)

func createProtection(org, repoName, protectionPattern string, minApprove int, dismissStalePrApprovals, codeOwner bool,
	requireBranchesUptodate, includeAdmins bool,
	canDismiss, canDismissTeams, canPush, canPushTeams []string) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	preq := &github.ProtectionRequest{
		RequiredPullRequestReviews: &github.PullRequestReviewsEnforcementRequest{
			RequireCodeOwnerReviews:      codeOwner,
			RequiredApprovingReviewCount: minApprove,
			DismissStaleReviews:          dismissStalePrApprovals,
		},
		RequiredStatusChecks: &github.RequiredStatusChecks{
			Strict:   requireBranchesUptodate,
			Contexts: []string{},
		},
		EnforceAdmins: includeAdmins,
	}

	if len(canDismiss)+len(canDismissTeams) > 0 {
		preq.RequiredPullRequestReviews.DismissalRestrictionsRequest = &github.DismissalRestrictionsRequest{
			Teams: &canDismissTeams,
			Users: &canDismiss,
		}
	}

	if len(canPush)+len(canPushTeams) > 0 {
		preq.Restrictions = &github.BranchRestrictionsRequest{
			Teams: canPushTeams,
			Users: canPush,
		}
	}

	_, _, err := client.Repositories.UpdateBranchProtection(ctx, org, repoName, protectionPattern, preq)
	if err == nil {
		fmt.Println(protectionPattern)
	} else {
		fmt.Println(err)
	}
}
