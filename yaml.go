package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/google/go-github/v28/github"
	"gopkg.in/yaml.v2"
)

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
type YamlOnCreate struct {
	AutoInit  *bool   `yaml:"autoInit"`
	Gitignore *string `yaml:"gitignore"`
	License   *string `yaml:"license"`
}
type YamlRepository struct {
	Name          *string `yaml:"name"`
	Description   *string `yaml:"description"`
	Homepage      *string `yaml:"homepage"`
	Private       *bool   `yaml:"private"`
	DefaultBranch *string `yaml:"defaultBranch"`

	OnCreate YamlOnCreate        `yaml:"onCreate"`
	Pages    YamlRepositoryPages `yaml:"pages"`
	Merge    YamlRepositoryMerge `yaml:"merge"`
	Teams    map[string]string   `yaml:"teams"`

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
type YamlBranchOnCreate struct {
	Commits []YamlBranchCommit `yaml:"commits"`
}
type YamlBranchCommit struct {
	Message     string `yaml:"message"`
	FileName    string `yaml:"fileName"`
	Text        string `yaml:"text"`
	Destination string `yaml:"destination"`
}
type YamlBranch struct {
	Name                 string                         `yaml:"name"`
	MinApprove           int                            `yaml:"minApprove"`
	CodeOwners           bool                           `yaml:"codeOwners"`
	IncludeAdmins        bool                           `yaml:"includeAdmins"`
	RequiredStatusChecks YamlBranchRequiredStatusChecks `yaml:"requiredStatusChecks"`
	Push                 YamlPush                       `yaml:"push"`
	OnCreate             YamlBranchOnCreate             `yaml:"onCreate"`
}

func repoToYaml(obj *github.Repository) YamlRepository {
	yamlRepo := YamlRepository{
		Name:          obj.Name,
		Description:   obj.Description,
		Homepage:      obj.Homepage,
		Private:       obj.Private,
		DefaultBranch: obj.DefaultBranch,

		OnCreate: YamlOnCreate{
			AutoInit:  obj.AutoInit,
			Gitignore: obj.GitignoreTemplate,
			License:   obj.LicenseTemplate,
		},
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

func applyYaml(org string, fileName string, format string) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	var yamlTop YamlTop
	if err := yaml.Unmarshal(bytes, &yamlTop); err != nil {
		fmt.Println("failed to parse ", fileName, ", err: ", err)
		return
	}

	fmt.Println("Parsed Yaml from file: ", yamlTop)

	repo := yamlTop.Github.Repository
	createOrUpdateRepo(org,
		repo.Name, repo.Description, repo.Homepage, repo.Private,
		not(repo.Pages.Issues), not(repo.Pages.Projects), not(repo.Pages.Wiki),
		repo.OnCreate.AutoInit, repo.OnCreate.Gitignore, repo.OnCreate.License,
		not(repo.Merge.AllowMergeCommit), not(repo.Merge.AllowRebaseMerge), not(repo.Merge.AllowSquashMerge),
		repo.DefaultBranch,
		format, true)
	for _, b := range repo.Branches {
		if b.Name != "master" {
			createBranch(org, *repo.Name, b.Name, format)
		}
		for _, c := range b.OnCreate.Commits {
			e := getPrimaryEmail()
			if c.FileName == "" && c.Text == "" {
				CheckIfError(fmt.Errorf("eighter specify fileName or text"))
			} else if c.Text != "" && c.Destination == "" {
				CheckIfError(fmt.Errorf("destination is empty"))
			} else if c.Text != "" && c.Destination != "" {
				c.FileName, err = createFileWithContent(c.Text)
				if err != nil {
					CheckIfError(fmt.Errorf("can't write to temp file"))
				}
			}

			addFile(org, *repo.Name, b.Name, c.FileName, c.Destination, c.Message, e, format)
		}
		createProtection(org, *repo.Name, b.Name, b.MinApprove,
			false, b.CodeOwners, b.RequiredStatusChecks.RequiredBranchesUpToDate, b.IncludeAdmins,
			"", "",
			makeCommaSeparatedString(b.Push.Users), makeCommaSeparatedString(b.Push.Teams), makeCommaSeparatedString(b.RequiredStatusChecks.Contexts))
	}
	for teamName, teamPerm := range repo.Teams {
		addTeamToRepo(org, *repo.Name, teamName, teamPerm)
	}
	createOrUpdateRepo(org,
		repo.Name, repo.Description, repo.Homepage, repo.Private,
		not(repo.Pages.Issues), not(repo.Pages.Projects), not(repo.Pages.Wiki),
		repo.OnCreate.AutoInit, repo.OnCreate.Gitignore, repo.OnCreate.License,
		not(repo.Merge.AllowMergeCommit), not(repo.Merge.AllowRebaseMerge), not(repo.Merge.AllowSquashMerge),
		repo.DefaultBranch,
		format, false)
}

func makeCommaSeparatedString(arr []string) string {
	var result string
	for _, v := range arr {
		result += v + ","
	}
	return result
}

func createFileWithContent(content string) (fileName string, err error) {
	fileName = os.TempDir() + TempFileName()
	Info(fileName)
	return fileName, ioutil.WriteFile(fileName, []byte(content), 0644)
}

// TempFileName generates a temporary filename for use in testing or whatever
func TempFileName() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return hex.EncodeToString(randBytes)
}
