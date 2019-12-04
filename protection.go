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
		preq.RequiredPullRequestReviews.DismissalRestrictionsRequest = &github.DismissalRestrictionsRequest{}
		if len(dismissUsers) > 0 {
			preq.RequiredPullRequestReviews.DismissalRestrictionsRequest.Users = &dismissUsers
		}
		if len(dismissTeams) > 0 {
			preq.RequiredPullRequestReviews.DismissalRestrictionsRequest.Teams = &dismissTeams
		}
	}

	pushUsers := splitArgs(canPush)
	pushTeams := splitArgs(canPushTeams)
	fmt.Println(canPushTeams, pushTeams)
	if len(pushUsers)+len(pushTeams) > 0 {
		preq.Restrictions = &github.BranchRestrictionsRequest{
			Users: []string{},
			Teams: []string{},
		}
		if len(pushUsers) > 0 {
			preq.Restrictions.Users = pushUsers
		}
		if len(pushTeams) > 0 {
			preq.Restrictions.Teams = pushTeams
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
	var result []string
	for _, s := range splitted {
		if len(strings.TrimSpace(s)) != 0 {
			result = append(result, s)
		}
	}
	return result
}
