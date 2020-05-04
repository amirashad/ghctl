package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/go-git/go-git/v5/storage/memory"
)

const gitUsername = "some-user"

// create branch from default branch
func createBranch(org string, repo string, branch string, format string) {
	auth := &http.BasicAuth{
		Username: gitUsername, // anything except an empty string
		Password: args.Token,
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

func addFile(org, repo, branch, fileFrom, fileTo, commitMessage, gitEmail, format string) {
	if fileFrom == "" {
		CheckIfError(fmt.Errorf("fileName is empty"))
	}
	if fileTo == "" {
		fileTo = fileFrom
	}
	if commitMessage == "" {
		commitMessage = "Change " + fileTo
	}
	auth := &http.BasicAuth{
		Username: gitUsername, // anything except an empty string
		Password: args.Token,
	}

	dir, err := ioutil.TempDir("", repo+"-"+branch)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	defer os.RemoveAll(dir)

	// Clone the given repository to the memory
	repoURL := fmt.Sprintf("https://github.com/%s/%s.git", org, repo)
	Info("git clone %s", repoURL)
	r, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:           repoURL,
		Auth:          auth,
		ReferenceName: plumbing.NewBranchReferenceName(branch),
		SingleBranch:  true,
		Depth:         1,
	})
	CheckIfError(err)

	w, err := r.Worktree()
	CheckIfError(err)

	// ... we need a file to commit so let's create a new file inside of the
	// worktree of the project using the go standard library.
	// Info("echo \"hello world!\" > example-git-file")
	// filename := filepath.Join(dir, files)
	copyFile(fileFrom, filepath.Join(dir, fileTo))
	// err = ioutil.WriteFile(filename, []byte("hello world!"), 0644)
	// CheckIfError(err)

	// Adds the new file to the staging area.
	Info("git add %s -> %s", fileFrom, fileTo)
	_, err = w.Add(fileTo)
	CheckIfError(err)

	// We can verify the current status of the worktree using the method Status.
	Info("git status --porcelain")
	status, err := w.Status()
	CheckIfError(err)

	fmt.Println(status)

	// Commits the current staging area to the repository, with the new file
	// just created. We should provide the object.Signature of Author of the
	// commit.
	Info("git commit -m \"%s\"", commitMessage)
	commit, err := w.Commit(commitMessage, &git.CommitOptions{
		Author: &object.Signature{
			Email: gitEmail,
			When:  time.Now(),
		},
	})

	CheckIfError(err)

	// Prints the current HEAD to verify that all worked well.
	Info("git show -s")
	obj, err := r.CommitObject(commit)
	CheckIfError(err)

	fmt.Println(obj)

	// push using default options
	Info("git push")
	err = r.Push(&git.PushOptions{Auth: auth})
	CheckIfError(err)
}

func ensureDirExists(fileName string) {
	dirName := filepath.Dir(fileName)
	os.MkdirAll(dirName, os.ModePerm)
}

func copyFile(from, to string) {
	input, err := ioutil.ReadFile(from)
	if err != nil {
		fmt.Println(err)
		return
	}

	ensureDirExists(to)
	err = ioutil.WriteFile(to, input, 0644)
	if err != nil {
		fmt.Println("Error creating", to)
		fmt.Println(err)
		return
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error) {
	if err == nil {
		return
	}

	Error(err)
	os.Exit(1)
}
