package git

import "github.com/libgit2/git2go"

type GitHead struct {
	workdir string
	repo    *git.Repository
}

func (head GitHead) Init(repo *git.Repository) *GitHead {
	head.workdir = repo.Workdir() + "/.git/"
	head.repo = repo
	return &head
}

func (head GitHead) IsRebasing() bool {
	reader := FileReader{}
	merge, err := reader.Exists(head.workdir + "rebase-merge")
	if nil != err {
		return false
	}

	interactive, err := reader.Exists(head.workdir + "rebase-apply")
	if nil != err {
		return false
	}

	return merge || interactive
}

func (head GitHead) IsMerging() bool {
	exists, err := FileReader{}.Exists(head.workdir + "MERGE_HEAD")
	if nil != err {
		return false
	}

	return exists
}

func (head GitHead) Name() (string, error) {
	if head.IsRebasing() {
		return head.rebasingBranchName()
	}

	if head.IsMerging() {
		return head.mergingBranchName()
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

func (head GitHead) mergingBranchName() (string, error) {
	name, err := head.branchName()
	if nil != err {
		return "", err
	}

	return "merging " + name, nil
}

func (head GitHead) rebasingBranchName() (string, error) {
	reference, err := head.reference(head.workdir + "rebase-apply/head-name")
	if nil != err {
		return "", err
	}

	if nil == reference {
		reference, err = head.reference(head.workdir + "rebase-merge/head-name")
		if nil != err {
			return "", err
		}
	}

	name, err := reference.Branch().Name()
	if nil != err {
		return "", err
	}

	reference.Free()
	return "rebasing " + name, nil
}

func (plugin GitHead) reference(path string) (*git.Reference, error) {
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

	return plugin.repo.LookupReference(head)
}
