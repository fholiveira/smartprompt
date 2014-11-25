package git

import "github.com/libgit2/git2go"

type GitChanges struct {
	workdir string
	repo    *git.Repository

	conflicted int
	staged     int
	modified   int
	untracked  int
}

func (changes GitChanges) Init(repo *git.Repository) (*GitChanges, error) {
	changes.workdir = repo.Workdir() + "/.git/"
	changes.repo = repo

	err := changes.countChanges()
	if nil != err {
		return nil, err
	}

	return &changes, nil
}

func (changes GitChanges) StagedFilesCount() int {
	return changes.staged
}

func (changes GitChanges) ModifiedFilesCount() int {
	return changes.modified
}

func (changes GitChanges) UntrackedFilesCount() int {
	return changes.untracked
}

func (changes GitChanges) ConflictedFilesCount() int {
	return changes.conflicted
}

func (changes GitChanges) HasChanges() bool {
	return changes.modified > 0 ||
		changes.staged > 0 ||
		changes.untracked > 0 ||
		changes.conflicted > 0
}

func (changes *GitChanges) checkForChanges(entry git.StatusEntry) {
	if isStaged(entry.Status) && isModified(entry.Status) {
		changes.conflicted++
	} else if isModified(entry.Status) {
		changes.modified++
	} else if isStaged(entry.Status) {
		changes.staged++
	} else if isUntracked(entry.Status) {
		changes.untracked++
	}
}

func (changes *GitChanges) countChanges() error {
	options := git.StatusOptions{}
	options.Flags = git.StatusOptIncludeUntracked

	status, err := changes.repo.StatusList(&options)
	if nil != err {
		return err
	}

	entryCount, err := status.EntryCount()
	if nil != err {
		return err
	}

	for index := 0; index < entryCount; index++ {
		entry, err := status.ByIndex(index)
		if nil != err {
			return err
		}

		changes.checkForChanges(entry)
	}

	return nil
}

func isStaged(status git.Status) bool {
	return status&git.StatusIndexNew > 0 ||
		status&git.StatusIndexModified > 0 ||
		status&git.StatusIndexDeleted > 0 ||
		status&git.StatusIndexRenamed > 0 ||
		status&git.StatusIndexTypeChange > 0
}

func isModified(status git.Status) bool {
	return status&git.StatusWtDeleted > 0 ||
		status&git.StatusWtModified > 0 ||
		status&git.StatusWtRenamed > 0 ||
		status&git.StatusWtTypeChange > 0
}

func isUntracked(status git.Status) bool {
	return status&git.StatusWtNew > 0
}
