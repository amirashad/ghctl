package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v28/github"
)

func createProtection(org, repoName, protectionPattern string, minApprove int, dismissStalePrApprovals, codeOwner bool,
	requireBranchesUptodate, includeAdmins bool,
	canDismiss, canDismissTeams, canPush, canPushTeams []string,
	requiredStatusChecks []string) {
	ctx := context.Background()
	client := authn(ctx)

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

	if len(requiredStatusChecks) > 0 {
		preq.RequiredStatusChecks.Contexts = requiredStatusChecks
	}

	_, _, err := client.Repositories.UpdateBranchProtection(ctx, org, repoName, protectionPattern, preq)
	if err == nil {
		fmt.Println(protectionPattern)
	} else {
		fmt.Println(err)
	}
}
