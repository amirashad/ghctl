package main

import (
	"fmt"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

const gitUsername = "some-user"

// create branch from default branch
func createBranch(org string, repo string, branch string, format string) {
	auth := &http.BasicAuth{
		Username: gitUsername, // anything except an empty string
		Password: githubToken(),
	}

	// Clone the given repository to the memory
	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", org, repo)
	Info("git clone %s", repoURL)
	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL:  repoURL,
		Auth: auth,
	})
	CheckIfError(err)

	// Create a new branch to the current HEAD
	Info("git branch %s", branch)
	headRef, err := r.Head()
	CheckIfError(err)
	fmt.Println(headRef)

	// Create a new plumbing.HashReference object with the name of the branch
	// and the hash from the HEAD. The reference name should be a full reference
	// name and not an abbreviated one, as is used on the git cli.
	ref := plumbing.NewHashReference(plumbing.NewBranchReferenceName(branch), headRef.Hash())

	// The created reference is saved in the storage.
	err = r.Storer.SetReference(ref)
	CheckIfError(err)

	// push using default options
	Info("git push")
	err = r.Push(&git.PushOptions{Auth: auth})
	CheckIfError(err)
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	Error(err)
	os.Exit(1)
}

// Info should be used to to display a info
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s\x1b[0m\n", fmt.Sprintf(format, args...))
}

// Error should be used to to display a error
func Error(err error, args ...interface{}) {
	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
}
