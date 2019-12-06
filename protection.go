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
	requiredStatusChecks string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

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

	dismissUsers := splitArgs(canDismiss)
	dismissTeams := splitArgs(canDismissTeams)
	if len(dismissUsers)+len(dismissTeams) > 0 {
		preq.RequiredPullRequestReviews.DismissalRestrictionsRequest = &github.DismissalRestrictionsRequest{
			Users: &dismissUsers,
			Teams: &dismissTeams,
		}
	}

	pushUsers := splitArgs(canPush)
	pushTeams := splitArgs(canPushTeams)
	if len(pushUsers)+len(pushTeams) > 0 {
		preq.Restrictions = &github.BranchRestrictionsRequest{
			Users: pushUsers,
			Teams: pushTeams, // TODO: teams not working
		}
	}

	if len(requiredStatusChecks) > 0 {
		preq.RequiredStatusChecks.Contexts = splitArgs(requiredStatusChecks)
	}

	_, _, err := client.Repositories.UpdateBranchProtection(ctx, org, repoName, protectionPattern, preq)
	if err == nil {
		fmt.Println(protectionPattern)
	} else {
		fmt.Println(err)
	}
}

func splitArgs(arg string) []string {
	splitted := strings.Split(arg, ",")
	result := []string{}
	for _, s := range splitted {
		if len(strings.TrimSpace(s)) != 0 {
			result = append(result, s)
		}
	}
	return result
}
