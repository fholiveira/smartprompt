package git

import "github.com/libgit2/git2go"

type RepoState git.RepositoryState

func (state RepoState) IsRebasing() bool {
	repostate := git.RepositoryState(state)
	return repostate == git.RepositoryStateRebaseInteractive
}

func (state RepoState) IsMerging() bool {
	repostate := git.RepositoryState(state)
	return repostate&git.RepositoryStateMerge > 0 ||
		repostate&git.RepositoryStateRebaseMerge > 0
}

func (state RepoState) IsReverting() bool {
	repostate := git.RepositoryState(state)
	return repostate&git.RepositoryStateRevert > 0
}

func (state RepoState) IsCherrypicking() bool {
	repostate := git.RepositoryState(state)
	return repostate&git.RepositoryStateCherrypick > 0
}

func (state RepoState) IsBisecting() bool {
	repostate := git.RepositoryState(state)
	return repostate&git.RepositoryStateBisect > 0
}

type GitHead struct {
	repo *git.Repository
}

func (head GitHead) Init(repo *git.Repository) *GitHead {
	head.repo = repo
	return &head
}

func (head GitHead) IsDetached() bool {
	detached, err := head.repo.IsHeadDetached()
	if nil != err {
		return false
	}

	return detached
}

func (head GitHead) Name() (string, error) {
	state := RepoState(head.repo.State())

	if state.IsRebasing() {
		name, err := head.branchName()
		return "rebasing " + name, err
	}

	if state.IsMerging() {
		name, err := head.branchName()
		return "merging " + name, err
	}

	if head.IsDetached() {
		reference, err := head.repo.Head()
		if nil != err {
			return "", err
		}

		return "detached at " + reference.Target().String()[0:7], nil
	}

	return head.branchName()
}

func (head GitHead) branchName() (string, error) {
	reference, err := head.repo.Head()
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
