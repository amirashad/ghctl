package main

import "github.com/google/go-github/v28/github"

type YamlRepositoryMerge struct {
	AllowMergeCommit *bool `yaml:"allowMergeCommit"`
	AllowSquashMerge *bool `yaml:"allowSquashMerge"`
	AllowRebaseMerge *bool `yaml:"allowRebaseMerge"`
}
type YamlRepositoryPages struct {
	Issues   *bool `yaml:"issues"`
	Projects *bool `yaml:"projects"`
	Wiki     *bool `yaml:"wiki"`
}
type YamlRepository struct {
	Name          *string `yaml:"name"`
	Description   *string `yaml:"description"`
	Homepage      *string `yaml:"homepage"`
	Private       *bool   `yaml:"private"`
	AutoInit      *bool   `yaml:"autoInit"`
	Gitignore     *string `yaml:"gitignore"`
	License       *string `yaml:"license"`
	DefaultBranch *string `yaml:"defaultBranch"`

	Pages YamlRepositoryPages `yaml:"pages"`
	Merge YamlRepositoryMerge `yaml:"merge"`
	Teams map[string]string   `yaml:"teams"`

	Branches []YamlBranch `yaml:"branches"`
}
type YamlGithub struct {
	Repository YamlRepository `yaml:"repo"`
}
type YamlTop struct {
	Github YamlGithub `yaml:"github"`
}

type YamlBranchRequiredStatusChecks struct {
	RequiredBranchesUpToDate bool     `yaml:"requiredBranchesUpToDate"`
	Contexts                 []string `yaml:"contexts"`
}
type YamlPush struct {
	Users []string `yaml:"users"`
	Teams []string `yaml:"teams"`
}
type YamlBranch struct {
	Name                 string                         `yaml:"name"`
	MinApprove           int                            `yaml:"minApprove"`
	CodeOwners           bool                           `yaml:"codeOwners"`
	IncludeAdmins        bool                           `yaml:"includeAdmins"`
	RequiredStatusChecks YamlBranchRequiredStatusChecks `yaml:"requiredStatusChecks"`
	Push                 YamlPush                       `yaml:"push"`
}

func repoToYaml(obj *github.Repository) YamlRepository {
	yamlRepo := YamlRepository{
		Name:          obj.Name,
		Description:   obj.Description,
		Homepage:      obj.Homepage,
		Private:       obj.Private,
		AutoInit:      obj.AutoInit,
		Gitignore:     obj.GitignoreTemplate,
		License:       obj.LicenseTemplate,
		DefaultBranch: obj.DefaultBranch,

		Pages: YamlRepositoryPages{
			Wiki:     obj.HasWiki,
			Projects: obj.HasProjects,
			Issues:   obj.HasIssues,
		},
		Merge: YamlRepositoryMerge{
			AllowMergeCommit: obj.AllowMergeCommit,
			AllowRebaseMerge: obj.AllowRebaseMerge,
			AllowSquashMerge: obj.AllowSquashMerge,
		},
		Teams:    getRepoTeams(*obj.Owner.Login, *obj.Name),
		Branches: getRepoProtections(*obj.Owner.Login, *obj.Name),
	}
	return yamlRepo
}
