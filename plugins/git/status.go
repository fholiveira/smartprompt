package git

import (
	"fmt"
	"os"
	"strings"

	"github.com/fholiveira/smartprompt/colors"
	"github.com/libgit2/git2go"
)

type GitStatus struct {
	repo *git.Repository
}

func (gitStatus GitStatus) repository() (*git.Repository, error) {
	currentDirectory, err := os.Getwd()
	if nil != err {
		return nil, err
	}

	repoDirectory, err := git.Discover(currentDirectory, false, nil)
	if err != nil {
		return nil, err
	}

	repo, err := git.OpenRepository(repoDirectory)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func (plugin GitStatus) textifyChanges(changes *GitChanges) (string, string) {
	if !changes.HasChanges() {
		return "", colors.Green.Bold()
	}

	if changes.ConflictedFilesCount() > 0 {
		return "!" + fmt.Sprint(changes.ConflictedFilesCount()), colors.Red.Bold()
	}

	prompt := fmt.Sprint(
		changes.StagedFilesCount(),
		changes.ModifiedFilesCount(),
		changes.UntrackedFilesCount())

	if changes.ModifiedFilesCount() > 0 {
		return prompt, colors.Red.Bold()
	}

	if changes.StagedFilesCount() > 0 {
		return prompt, colors.Yellow.Bold()
	}

	return prompt, colors.Cyan.Bold()
}

func (gitStatus GitStatus) IsApplicable() bool {
	repo, _ := gitStatus.repository()
	return nil != repo
}

func (plugin GitStatus) Prompt(parameters []string) (string, error) {
	repo, err := plugin.repository()
	if nil != err {
		return "", nil
	}

	changes, err := GitChanges{}.Init(repo)
	if nil != err {
		return "", nil
	}

	name, err := GitHead{}.Init(repo).Name()
	changesText, color := plugin.textifyChanges(changes)
	prompt := strings.TrimSpace(name + " " + changesText)

	repo.Free()

	return color + "[" + prompt + "]", nil
}
