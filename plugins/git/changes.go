package git

import "github.com/libgit2/git2go"

type GitChanges struct {
	workdir string
	repo    *git.Repository

	staged    int
	modified  int
	untracked int
}

func (changes GitChanges) Init(repo *git.Repository) *GitChanges {
	changes.workdir = repo.Workdir() + "/.git/"
	changes.repo = repo

	changes.countChanges()

	return &changes
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

func (changes GitChanges) HasChanges() bool {
	return changes.modified > 0 || changes.staged > 0 || changes.untracked > 0
}

func (changes *GitChanges) countChanges() {
	options := git.StatusOptions{}
	options.Flags = git.StatusOptIncludeUntracked

	status, _ := changes.repo.StatusList(&options)
	entryCount, _ := status.EntryCount()

	for index := 0; index < entryCount; index++ {
		entry, _ := status.ByIndex(index)

		if isStaged(entry.Status) {
			changes.staged++
		} else if isModified(entry.Status) {
			changes.modified++
		} else if isUntracked(entry.Status) {
			changes.untracked++
		}
	}
}

func isStaged(status git.Status) bool {
	return status&git.StatusIndexNew > 0 ||
		status&git.StatusIndexModified > 0 ||
		status&git.StatusWtDeleted > 0 ||
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
