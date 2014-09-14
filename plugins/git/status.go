package git

import (
	"fmt"
	"os"

	"github.com/libgit2/git2go"
)

type GitStatus struct{}

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

	head := GitHead{}.Init(repo)
	prompt, err := head.Name()
	if nil != err {
		return "", err
	}

	changes := GitChanges{}.Init(repo)
	if changes.HasChanges() {
		if changes.ConflictedFilesCount() > 0 {
			prompt += " !" + fmt.Sprint(changes.ConflictedFilesCount())
			color = "{RED:bold}"
		} else {
			prompt += " " + fmt.Sprint(changes.StagedFilesCount(), changes.ModifiedFilesCount(), changes.UntrackedFilesCount())

			if changes.ModifiedFilesCount() > 0 {
				color = "{RED:bold}"
			} else if changes.StagedFilesCount() > 0 {
				color = "{YELLOW:bold}"
			} else if changes.UntrackedFilesCount() > 0 {
				color = "{CYAN:bold}"
			}
		}
	}

	repo.Free()

	return color + "[" + prompt + "]", nil
}
