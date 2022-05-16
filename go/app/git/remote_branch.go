package git

import (
	"github.com/go-git/go-git/v5/plumbing"
	"strings"
	"time"
)

type RemoteBranch struct {
	RemoteName string
	Parent     *plumbing.Reference
	Repository *Repository
}

func (r *Remote) GetBranches(filter string) ([]*RemoteBranch, error) {

	references, err := r.Repository.Parent.References()
	if err != nil {
		return nil, err
	}

	var remoteBranches []*RemoteBranch
	err = references.ForEach(func(ref *plumbing.Reference) error {
		if r.IsParentOfBranch(ref) {

			newBranch := RemoteBranch{
				Parent:     ref,
				Repository: r.Repository,
				RemoteName: r.Name(),
			}

			if newBranch.IsFilterPassed(filter) {
				remoteBranches = append(remoteBranches, &newBranch)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return remoteBranches, nil
}

func (rb *RemoteBranch) ShortName() string {
	shortName := strings.TrimPrefix(rb.Parent.Name().Short(), rb.RemoteName+"/")
	return shortName
}

func (rb *RemoteBranch) Hash() plumbing.Hash {
	return rb.Parent.Hash()
}

func (rb *RemoteBranch) LastCommit() (*Commit, error) {
	return rb.Repository.GetLastCommit(rb.Hash())
}

func (rb *RemoteBranch) LastCommitAuthor(def string) string {
	lastCommit, err := rb.LastCommit()
	if err != nil {
		return def
	}

	return lastCommit.Author()
}

func (rb *RemoteBranch) LastCommitTime(def time.Time) time.Time {
	lastCommit, err := rb.LastCommit()
	if err != nil {
		return def
	}

	return lastCommit.Time()
}

func (rb *RemoteBranch) ToRow() []string {

	lastCommit, err := rb.LastCommit()
	if err != nil {
		lastCommit = &Commit{}
	}

	return []string{rb.ShortName(), lastCommit.Time().Format("2006-01-02 15:04-07"), lastCommit.Author(), lastCommit.Title()}
}

func (rb *RemoteBranch) IsFilterPassed(filter string) bool {

	if filter == "" {
		return true
	}

	lastCommit, err := rb.LastCommit()
	if err != nil {
		return false
	}

	return strings.Contains(rb.ShortName(), filter) ||
		strings.Contains(lastCommit.Title(), filter) ||
		strings.Contains(lastCommit.Author(), filter)
}

func RemoteBranchHeaders() []string {
	return []string{"Branch", "Date", "Author", "Title"}
}
