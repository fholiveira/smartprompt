package plugins

import (
	"fmt"

	"github.com/libgit2/git2go"
)
import "os"

type Git struct{}

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

func (git Git) Prompt() (string, error) {
	repo, err := getRepository()
	if nil != err {
		return "", err
	}

	branchName, err := getBranchName(repo)
	if nil != err {
		return "", err
	}

	index, err := repo.Index()
	if nil != err {
		return "", err
	}

	fmt.Println(index.HasConflicts())

	return "[" + branchName + "]", nil
}
