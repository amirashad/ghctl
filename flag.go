package main

type Args struct {
	Token        string `arg:"env:GITHUB_TOKEN,required"`
	Org          string `arg:"env:GITHUB_ORG,required"`
	OutputFormat string `arg:"-o" help:"output format: normal, json" default:"normal"`
	Verbose      bool   `arg:"-v" default:"false"`

	Get    *Get    `arg:"subcommand:get"`
	Create *Create `arg:"subcommand:create"`
	Add    *Add    `arg:"subcommand:add"`
	Update *Create `arg:"subcommand:update"`
	Apply  *Apply  `arg:"subcommand:apply"`
}

func (Args) Version() string {
	return "0.3.2"
}

type Get struct {
	Repos   *Repos   `arg:"subcommand:repos"`
	Members *Members `arg:"subcommand:members"`
	Teams   *Teams   `arg:"subcommand:teams"`
}
type Repos struct {
	RepoName *string `arg:"positional"`
}
type Members struct {
}
type Teams struct {
	TeamName *string `arg:"positional"`
}

type Create struct {
	Repo       *Repo       `arg:"subcommand:repo"`
	Branch     *Branch     `arg:"subcommand:branch"`
	Protection *Protection `arg:"subcommand:protection"`
}
type Repo struct {
	Name        *string `arg:"-n,required"`
	Description *string `arg:"-d"`
	Homepage    *string `arg:"-h"`

	Private    *bool
	NoIssues   *bool
	NoProjects *bool
	NoWiki     *bool

	AutoInit          *bool   `arg:"-a"`
	GitignoreTemplate *string `arg:"-i"`
	LicenseTemplate   *string `arg:"-l"`

	NoMergeCommit *bool
	NoSquashMerge *bool
	NoRebaseMerge *bool

	DefaultBranch *string
}
type Branch struct {
	Repo   string `arg:"-r,required"`
	Branch string `arg:"-b,required"`
}
type Protection struct {
	Repo   string `arg:"-r,required"`
	Branch string `arg:"-b,required"`

	MinApprove              int  `arg:"-p"`
	DismissStaleReviews     bool `arg:"-d"`
	CanDismiss              string
	CanDismissTeams         string
	RequireBranchesUpToDate bool

	CodeOwner     bool `arg:"-c"`
	IncludeAdmins bool `arg:"-a"`

	CanPush      string
	CanPushTeams string

	RequiredStatusChecks string `arg:"-s,--required-status-checks"`
}

type Add struct {
	File         *File         `arg:"subcommand:file"`
	Collaborator *Collaborator `arg:"subcommand:collaborator"`
	Team         *Team         `arg:"subcommand:team"`
}
type File struct {
	Repo          string `arg:"-r,required"`
	Branch        string `arg:"-b,required"`
	File          string `arg:"-f,required"`
	GitName       string `arg:"-n,required"`
	GitEmail      string `arg:"-e,required"`
	CommitMessage string `arg:"-m,--gitmessage"`
}
type Collaborator struct {
	Repo       string `arg:"-r,required"`
	User       string `arg:"-u,required"`
	Permission string `arg:"-p,required"`
}
type Team struct {
	Repo       string `arg:"-r,required"`
	Team       string `arg:"-t,required"`
	Permission string `arg:"-p,required"`
}

type Apply struct {
	FileName string `arg:"-f,required"`
}
