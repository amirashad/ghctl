# Change log

## v0.5.1 (2021-04-05)

 - Fix the whitespaces issues when reading yaml list

## v0.5.0 (2021-03-04)

 - Update go version to 1.16, github-go version to v33, go-git version to v5(.2)

## v0.4.3 (2020-05-04)

 - Add new arg for repositories to delete branch on merge: `deleteBranchOnMerge`
 - Update go version to 1.14, github-go version to v31, go-git version to v5

## v0.4.1 (2020-03-05)

 - No need `--gitname` param for `add file` command anymore
 - FIX: Now destination filename is source filename, instead of commitmessage

## v0.4.0 (2020-03-05)

 - Update go-github dependency to latest (v29)

## v0.3.5 (2020-02-11)

 - FIX: YAML: Add teams to repo before adding users to repo as mergers

## v0.3.4 (2019-12-17)

 - README: Add installation instruction for Windows
 - README: Enhance features

## v0.3.3 (2019-12-17)

 - Add example for repo-creation
 - Add demo gif

## v0.3.2 (2019-12-13)

 - Update readme to make it understandable

## v0.3.1 (2019-12-09)

 - Docker: Add yq tool to docker image

## v0.3.0 (2019-12-07)

 - Docker: Create DockerHub image

## v0.2.9 (2019-12-06)

 - YAML: Add support to get file content from yaml

## v0.2.8 (2019-12-06)

 - YAML: Add files to branch with commits on creation

## v0.2.7 (2019-12-04)

 - Support for `apply` command to apply yaml. Now creates repo, adds teams to repo, creates branches and protections

## v0.2.6 (2019-12-02)

 - Support for `apply` command to apply yaml then

## v0.2.5 (2019-12-02)

 - Add `onCreate` tag to yaml and group them (autoInit, gitignore, license)

## v0.2.4.1 (2019-12-02)

 - Comment SNYK Badge, because of they don't support Go Github projects yet. 

## v0.2.4 (2019-12-01)

 - Get repo protection information as yaml: `get repo REPONAME -o yaml`

## v0.2.3 (2019-11-25)

 - Get repo detailed information by name: `get repo REPONAME -o json|yaml`
 - Get repo as yaml: `get repo REPONAME -o yaml`

## v0.2.2 (2019-11-25)

 - Change multiple argument strategy. Now we should add `,` between args of canpush, canpushteams, candismiss, candismissteams, required-status-checks

## v0.2.1 (2019-11-11)

 - Add team to repo: `add team --team "team" --repo "repo" --permission "pull|push|admin"`

## v0.2.0 (2019-11-11)

 - Get detailed information about team by team name: `get teams TEAMNAME`

## v0.1.10 (2019-11-09)

 - Update `create repo` command with support for `--defaultbranch "develop"`
 - Create `update repo` command analogy with `create repo`

## v0.1.9 (2019-11-08)

 - Add go style checking (golint) and vet to pipeline
 - Add SNYK vulnerability check to pipeline
 - Add SNYK badge to README

## v0.1.8 (2019-11-08)

 - Accept org name from env variable: GITHUB_ORG

## v0.1.7 (2019-11-01)

 - Update go-arg to version 1.2.0
 - Add some unit tests

## v0.1.6.1 (2019-11-01)

 - Change github badge style to default rounded rect

## v0.1.6 (2019-10-17)

 - Add `add collaborator` command with flags `--org SomeOrg --repo some-repo --user some-user --permission pull|push|admin`

## v0.1.5 (2019-10-15)

 - Update `create protection` command with flags `-s|--required-status-checks "ci/circleci: build" "SonarCloud Code Analysis"`

## v0.1.4 (2019-10-15)

 - Change CircleCI status badge type
 - Add SonarCloud properties to check bugs and keep code clean
 - Add SonarCloud status badge
 - Create GitHub releases only with version

## v0.1.3 (2019-10-14)

 - Add Github release badge

## v0.1.2 (2019-10-14)

 - Add CircleCI status badge

## v0.1.1 (2019-10-12)

 - Use struct based flags
 - Add Core Infrastructure Initiative (CII) badge
 - Delete wide from output types
 - Change command line arguments

## v0.1.0 (2019-10-12)

 - Update `create protection` command with flags `-can-push "user1,user2" -can-push-teams "team1,team2" -can-dismiss "user1,user2" -can-dismiss-teams "team1,team2"`

## v0.0.9 (2019-10-10)

 - Update `create protection` command with flags `-require-branches-uptodate true|false -admins true|false`

## v0.0.8 (2019-10-09)

 - Add `create protection` command with flags `-repo reponame -p protection-pattern -min-approve count -dismiss-stale-pr-approvals true|false -code-owner true|false`

## v0.0.7 (2019-10-07)

 - Add `add file` command with flags `-repo reponame -b branchname -f file -gitname "Author Name" -gitemail "author.email@email.com" -m "Commit message"`

## v0.0.6 (2019-10-05)

 - Add `create branch` command with flags `-repo reponame -b branchname`

## v0.0.5 (2019-10-05)

 - Support `create repo` command with flags `-n name -d description -h homepage`
 - Support `create repo` command with flags `-private true|false -issues true|false -projects true|false -wikis true|false`
 - Support `create repo` command with flags `-a true|false -g gitignoretemplate -l license`
 - Support `create repo` command with flags `-mergecommit true|false -squash true|false -rebase true|false`

## v0.0.4 (2019-10-04)

 - Add `create repo -n "repo-name" -o [normal, json, wide]` command

## v0.0.3 (2019-10-04)

 - Add `get teams` command
 - Add `get members -o [normal, json, wide]` command

## v0.0.2 (2019-10-02)

 - Add `get members` command
 - Add `get members -o [normal, json, wide]` command

## v0.0.1 (2019-09-30)

 - Add `get repos` command
 - Add `get repos -o [normal, json, wide]` command
