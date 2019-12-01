package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/google/go-github/v28/github"
	"gopkg.in/yaml.v2"
)

func getRepos(org string, format string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	opt := &github.RepositoryListByOrgOptions{ /*Type: "private", */ ListOptions: github.ListOptions{PerPage: 100}}
	var objsAll []*github.Repository
	for {
		objs, resp, err := client.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		objsAll = append(objsAll, objs...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	sort.Slice(objsAll, func(i, j int) bool {
		return *objsAll[i].Name < *objsAll[j].Name
	})

	if format == "normal" {
		for _, repo := range objsAll {
			fmt.Println(*repo.Name)
		}
	} else if format == "json" {
		bytes, _ := json.Marshal(objsAll)
		fmt.Println(string(bytes))
	}
}

func createOrUpdateRepo(org string,
	name, descr, homepage *string,
	private, noIssues, noProjects, noWiki, autoinit *bool,
	gitIgnoreTemplate, licenseTemplate *string,
	noMergeCommit, noSquashMerge, noRebaseMerge *bool,
	defaultBranch *string,
	format string, create bool) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	repo := &github.Repository{
		Name:        name,
		Description: descr,
		Homepage:    homepage,

		Private:     private,
		HasIssues:   not(noIssues),
		HasProjects: not(noProjects),
		HasWiki:     not(noWiki),
		AutoInit:    autoinit,

		GitignoreTemplate: gitIgnoreTemplate,
		LicenseTemplate:   licenseTemplate,

		AllowMergeCommit: not(noMergeCommit),
		AllowSquashMerge: not(noSquashMerge),
		AllowRebaseMerge: not(noRebaseMerge),

		DefaultBranch: defaultBranch,
	}

	var objs *github.Repository
	var err error
	if create {
		objs, _, err = client.Repositories.Create(ctx, org, repo)
	} else {
		objs, _, err = client.Repositories.Edit(ctx, org, *name, repo)
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if format == "normal" {
		fmt.Println(*objs.Name)
	} else if format == "json" {
		bytes, _ := json.Marshal(objs)
		fmt.Println(string(bytes))
	}
}

func not(o *bool) *bool {
	result := true
	if o == nil {
		return &result
	}
	result = !*o
	return &result
}

func addCollaboratorToRepo(org string,
	repo, user, permission string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	perm := &github.RepositoryAddCollaboratorOptions{
		Permission: permission,
	}

	resp, err := client.Repositories.AddCollaborator(ctx, org, repo, user, perm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp.Status)
}

func getRepo(org string, repo *string, format string) {
	ctx := context.Background()
	client := createGithubClient(ctx)

	obj, _, err := client.Repositories.Get(ctx, org, *repo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if format == "normal" {
		fmt.Println(*obj.Name)
	} else if format == "json" {
		bytes, _ := json.Marshal(obj)
		fmt.Println(string(bytes))
	} else if format == "yaml" {
		yamlTop := YamlTop{
			Github: YamlGithub{
				Repository: repoToYaml(obj),
			},
		}
		bytes, err := yaml.Marshal(&yamlTop)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(bytes))
	}
}

func getRepoTeams(org string, repo string) map[string]string {
	ctx := context.Background()
	client := createGithubClient(ctx)

	teams, _, err := client.Repositories.ListTeams(ctx, org, repo, &github.ListOptions{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mapTeams := map[string]string{}
	for _, v := range teams {
		mapTeams[*v.Name] = *v.Permission
	}

	return mapTeams
}

func getRepoProtections(org string, repo string) []YamlBranch {
	ctx := context.Background()
	client := createGithubClient(ctx)

	opt := &github.ListOptions{PerPage: 100}
	var objsAll []*github.Branch
	for {
		branches, resp, err := client.Repositories.ListBranches(ctx, org, repo, opt)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		objsAll = append(objsAll, branches...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	protectedBranches := []YamlBranch{}
	for _, v := range objsAll {
		if *v.Protected {
			protection, _, err := client.Repositories.GetBranchProtection(ctx, org, repo, *v.Name)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			var users, teams []string
			if protection.Restrictions != nil {
				for _, v := range protection.Restrictions.Users {
					users = append(users, *v.Login)
				}
				for _, v := range protection.Restrictions.Teams {
					teams = append(teams, *v.Name)
				}
			}

			yamlBranch := YamlBranch{
				Name:          *v.Name,
				MinApprove:    protection.RequiredPullRequestReviews.RequiredApprovingReviewCount,
				CodeOwners:    protection.RequiredPullRequestReviews.RequireCodeOwnerReviews,
				IncludeAdmins: protection.EnforceAdmins.Enabled,
				RequiredStatusChecks: YamlBranchRequiredStatusChecks{
					RequiredBranchesUpToDate: protection.RequiredStatusChecks.Strict,
					Contexts:                 protection.RequiredStatusChecks.Contexts,
				},
				Push: YamlPush{
					Users: users,
					Teams: teams,
				},
			}
			protectedBranches = append(protectedBranches, yamlBranch)
		}
	}

	return protectedBranches
}
