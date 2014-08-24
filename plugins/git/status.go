package git

import (
	"os"

	"github.com/fholiveira/git2go"
)

type GitStatus struct{}

func getBranchName(repo *git.Repository) (string, error) {
	reference, err := repo.Head()
	if nil != err {
		return "", err
	}

	return reference.Branch().Name()
}

func getRepository() (*git.Repository, error) {
	workingDirectory, err := os.Getwd()
	if nil != err {
		return nil, err
	}

	repo, err := git.OpenRepository(workingDirectory)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (git GitStatus) Prompt(parameter string) (string, error) {
	repo, err := getRepository()
	if nil != err {
		return "", nil
	}

	rebase := GitRebase{}.Init(repo)
	if rebase.IsRebasing() {
		status, err := rebase.Status()
		if nil != err {
			return "", nil
		}

		return "{RED:bold}[" + status + "]", nil
	}

	branchName, err := getBranchName(repo)
	if nil != err {
		return "", err
	}

	return "{GREEN:bold}[" + branchName + "]", nil
}
