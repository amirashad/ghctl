package main

import (
	"github.com/alexflint/go-arg"
)

var args Args

func main() {
	args.OutputFormat = "normal"
	arg.MustParse(&args)
	// fmt.Println(args)

	if args.Get != nil && args.Get.Repos != nil {
		getRepos(args.Org, args.OutputFormat)
	} else if args.Get != nil && args.Get.Members != nil {
		getMembers(args.Org, args.OutputFormat)
	} else if args.Get != nil && args.Get.Teams != nil {
		getTeams(args.Org, args.OutputFormat)
	} else if args.Create != nil && args.Create.Repo != nil {
		createRepo(args.Org,
			args.Create.Repo.Name, args.Create.Repo.Description, args.Create.Repo.Homepage,
			args.Create.Repo.Private, args.Create.Repo.NoIssues, args.Create.Repo.NoProjects, args.Create.Repo.NoWiki, args.Create.Repo.AutoInit,
			args.Create.Repo.GitignoreTemplate, args.Create.Repo.LicenseTemplate,
			args.Create.Repo.NoMergeCommit, args.Create.Repo.NoSquashMerge, args.Create.Repo.NoRebaseMerge,
			args.OutputFormat)
	} else if args.Create != nil && args.Create.Branch != nil {
		createBranch(args.Org,
			args.Create.Branch.Repo,
			args.Create.Branch.Branch,
			args.OutputFormat)
	} else if args.Add != nil && args.Add.Files != nil {
		addFiles(args.Org,
			args.Add.Files.Repo, args.Add.Files.Branch,
			args.Add.Files.Files,
			args.Add.Files.CommitMessage,
			args.Add.Files.GitName, args.Add.Files.GitEmail,
			args.OutputFormat)
	} else if args.Create != nil && args.Create.Protection != nil {
		createProtection(args.Org,
			args.Create.Protection.Repo, args.Create.Protection.Branch,
			args.Create.Protection.MinApprove,
			args.Create.Protection.DismissStaleReviews,
			args.Create.Protection.CodeOwner,
			args.Create.Protection.RequireBranchesUpToDate,
			args.Create.Protection.IncludeAdmins,
			args.Create.Protection.CanDismiss, args.Create.Protection.CanDismissTeams,
			args.Create.Protection.CanPush, args.Create.Protection.CanPushTeams,
			args.Create.Protection.RequiredStatusChecks)
	}
}

func githubToken() string {
	return args.Token
}
