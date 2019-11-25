package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v28/github"
)

func createProtection(org, repoName, protectionPattern string, minApprove int, dismissStalePrApprovals, codeOwner bool,
	requireBranchesUptodate, includeAdmins bool,
	canDismiss, canDismissTeams, canPush, canPushTeams string,
	requiredStatusChecks string,
	add, delete bool) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	// protection, _, err := client.Repositories.GetBranchProtection(ctx, org, repoName, protectionPattern)
	// if err != nil {

	// }

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
		teams := strings.Split(canDismissTeams, ",")
		users := strings.Split(canDismiss, ",")
		preq.RequiredPullRequestReviews.DismissalRestrictionsRequest = &github.DismissalRestrictionsRequest{
			Teams: &teams,
			Users: &users,
		}
	}

	if len(canPush)+len(canPushTeams) > 0 {
		preq.Restrictions = &github.BranchRestrictionsRequest{
			Teams: strings.Split(canPushTeams, ","),
			Users: strings.Split(canPush, ","),
		}
	}

	if len(requiredStatusChecks) > 0 {
		preq.RequiredStatusChecks.Contexts = strings.Split(requiredStatusChecks, ",")
	}

	_, _, err := client.Repositories.UpdateBranchProtection(ctx, org, repoName, protectionPattern, preq)
	if err == nil {
		fmt.Println(protectionPattern)
	} else {
		fmt.Println(err)
	}
}
