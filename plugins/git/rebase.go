package git

import "github.com/libgit2/git2go"

type GitRebase struct {
	workdir string
	repo    *git.Repository
}

func (rebase GitRebase) Init(repo *git.Repository) *GitRebase {
	rebase.workdir = repo.Workdir() + "/.git/"
	rebase.repo = repo
	return &rebase
}

func (rebase GitRebase) IsRebasing() bool {
	reader := FileReader{}
	merge, err := reader.Exists(rebase.workdir + "rebase-merge")
	if nil != err {
		return false
	}

	interactive, err := reader.Exists(rebase.workdir + "rebase-apply")
	if nil != err {
		return false
	}

	return merge || interactive
}

func (rebase GitRebase) Status() (string, error) {
	ref, err := rebase.tryGetRef(rebase.workdir + "rebase-apply/head-name")
	if nil != err {
		return "", err
	}

	if nil == ref {
		ref, err = rebase.tryGetRef(rebase.workdir + "rebase-merge/head-name")
		if nil != err {
			return "", err
		}
	}

	name, err := ref.Branch().Name()
	if nil != err {
		return "", err
	}

	ref.Free()
	return "rebasing " + name, nil
}

func (rebase GitRebase) tryGetRef(path string) (*git.Reference, error) {
	reader := FileReader{}

	exists, err := reader.Exists(path)
	if nil != err {
		return nil, err
	}

	if !exists {
		return nil, nil
	}

	head, err := reader.ReadFirstLine(path)
	if nil != err {
		return nil, err
	}

	return rebase.repo.LookupReference(head)
}
