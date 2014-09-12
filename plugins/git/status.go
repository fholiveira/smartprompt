package git

import (
	"fmt"
	"os"

	"github.com/libgit2/git2go"
)

type GitStatus struct{}

func getBranchName(repo *git.Repository) (string, error) {
	reference, err := repo.Head()
	if nil != err {
		return "", err
	}

	name, err := reference.Branch().Name()
	if nil != err {
		return "", err
	}

	reference.Free()
	return name, nil
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

func (gitStatus GitStatus) Prompt(parameter string) (string, error) {

	repo, err := getRepository()
	if nil != err {
		return "", nil
	}

	color := "{GREEN:bold}"
	prompt, err := getBranchName(repo)
	if nil != err {
		return "", err
	}

	rebase := GitRebase{}.Init(repo)
	if rebase.IsRebasing() {
		prompt, err = rebase.Status()
		if nil != err {
			return "", nil
		}

		color = "{RED:bold}"
	}

	changes := GitChanges{}.Init(repo)
	if changes.HasChanges() {
		prompt += " " + fmt.Sprint(changes.StagedFilesCount(), changes.ModifiedFilesCount(), changes.UntrackedFilesCount())

		if changes.ModifiedFilesCount() > 0 {
			color = "{RED:bold}"
		} else if changes.StagedFilesCount() > 0 {
			color = "{YELLOW:bold}"
		} else if changes.UntrackedFilesCount() > 0 {
			color = "{CYAN:bold}"
		}
	}

	repo.Free()

	return color + "[" + prompt + "]", nil
}
