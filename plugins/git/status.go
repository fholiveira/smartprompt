package git

import (
	"fmt"
	"os"
	"strings"

	"github.com/libgit2/git2go"
)

type GitStatus struct {
	repo *git.Repository
}

func (gitStatus *GitStatus) repository() (*git.Repository, error) {
	if nil != gitStatus.repo {
		return gitStatus.repo, nil
	}

	workingDirectory, err := os.Getwd()
	if nil != err {
		return nil, err
	}

	repo, err := git.OpenRepository(workingDirectory)
	if err != nil {
		return nil, err
	}

	gitStatus.repo = repo
	return repo, nil
}

func (gitStatus GitStatus) IsApplicable() bool {
	repo, _ := gitStatus.repository()
	return nil != repo
}

func (gitStatus GitStatus) Prompt(parameter string) (string, error) {
	repo, err := gitStatus.repository()
	if nil != err {
		return "", nil
	}
	color := "{GREEN:bold}"

	head := GitHead{}.Init(repo)
	prompt, _ := head.Name()

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

	prompt = strings.TrimSpace(prompt)
	return color + "[" + prompt + "]", nil
}
