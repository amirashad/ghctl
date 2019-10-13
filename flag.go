package main

type Args struct {
	Token        string `arg:"env:GITHUB_TOKEN,required"`
	Org          string `arg:"required"`
	OutputFormat string `arg:"-o" help:"output format: normal, json"`

	Get    *Get    `arg:"subcommand:get"`
	Create *Create `arg:"subcommand:create"`
	Add    *Add    `arg:"subcommand:add"`
}

func (Args) Version() string {
	return "v0.1.1"
}

type Get struct {
	Repos   *Repos   `arg:"subcommand:repos"`
	Members *Members `arg:"subcommand:members"`
	Teams   *Teams   `arg:"subcommand:teams"`
}
type Repos struct {
}
type Members struct {
}
type Teams struct {
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
	CanDismiss              []string
	CanDismissTeams         []string
	RequireBranchesUpToDate bool

	CodeOwner     bool `arg:"-c"`
	IncludeAdmins bool `arg:"-a"`

	CanPush      []string
	CanPushTeams []string
}

type Add struct {
	Files *Files `arg:"subcommand:file"`
}
type Files struct {
	Repo          string   `arg:"-r,required"`
	Branch        string   `arg:"-b,required"`
	Files         []string `arg:"-f,required"`
	GitName       string   `arg:"-n,required"`
	GitEmail      string   `arg:"-e,required"`
	CommitMessage string   `arg:"-m"`
}
