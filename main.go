package main

import (
	"github.com/alexflint/go-arg"
)

var args Args

func main() {
	arg.MustParse(&args)
	// fmt.Println(args)

	if args.Get != nil && args.Get.Repos != nil {
		if args.Get.Repos.RepoName != nil {
			getRepo(args.Org, args.Get.Repos.RepoName, args.OutputFormat)
		} else {
			getRepos(args.Org, args.OutputFormat)
		}
	} else if args.Get != nil && args.Get.Members != nil {
		getMembers(args.Org, args.OutputFormat)
	} else if args.Get != nil && args.Get.Teams != nil {
		if args.Get.Teams.TeamName != nil {
			getTeam(args.Org, args.Get.Teams.TeamName, args.OutputFormat)
		} else {
			getTeams(args.Org, args.OutputFormat)
		}
	} else if args.Create != nil && args.Create.Repo != nil {
		createOrUpdateRepo(args.Org,
			args.Create.Repo.Name, args.Create.Repo.Description, args.Create.Repo.Homepage,
			args.Create.Repo.Private, args.Create.Repo.NoIssues, args.Create.Repo.NoProjects, args.Create.Repo.NoWiki, args.Create.Repo.AutoInit,
			args.Create.Repo.GitignoreTemplate, args.Create.Repo.LicenseTemplate,
			args.Create.Repo.NoMergeCommit, args.Create.Repo.NoSquashMerge, args.Create.Repo.NoRebaseMerge,
			args.Create.Repo.DefaultBranch,
			args.OutputFormat, true)
	} else if args.Update != nil && args.Update.Repo != nil {
		createOrUpdateRepo(args.Org,
			args.Update.Repo.Name, args.Update.Repo.Description, args.Update.Repo.Homepage,
			args.Update.Repo.Private, args.Update.Repo.NoIssues, args.Update.Repo.NoProjects, args.Update.Repo.NoWiki, args.Update.Repo.AutoInit,
			args.Update.Repo.GitignoreTemplate, args.Update.Repo.LicenseTemplate,
			args.Update.Repo.NoMergeCommit, args.Update.Repo.NoSquashMerge, args.Update.Repo.NoRebaseMerge,
			args.Update.Repo.DefaultBranch,
			args.OutputFormat, false)
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
	} else if args.Add != nil && args.Add.Collaborator != nil {
		addCollaboratorToRepo(args.Org,
			args.Add.Collaborator.Repo, args.Add.Collaborator.User,
			args.Add.Collaborator.Permission)
	} else if args.Add != nil && args.Add.Team != nil {
		addTeamToRepo(args.Org,
			args.Add.Team.Repo, args.Add.Team.Team,
			args.Add.Team.Permission)
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
	} else if args.Apply != nil {
		applyYaml(args.Org, args.Apply.FileName, args.OutputFormat)
	}
}
