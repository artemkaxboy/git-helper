package git

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"strings"
)

type Remote struct {
	Repository *Repository
	Parent     *git.Remote
}

func (r *Repository) GetRemotes() ([]*Remote, error) {
	nativeRemotes, err := r.Parent.Remotes()
	if err != nil {
		return nil, err
	}

	remotes := make([]*Remote, len(nativeRemotes))
	for i, remote := range nativeRemotes {
		remotes[i] = &Remote{Repository: r, Parent: remote}
	}

	return remotes, nil
}

func (r *Remote) Name() string {
	return r.Parent.Config().Name
}

func (r *Remote) GetRefPrefix() string {
	return r.Name() + "/"
}

func (r *Remote) IsParentOfBranch(ref *plumbing.Reference) bool {
	return ref.Name().IsRemote() &&
		ref.Type() != plumbing.SymbolicReference &&
		strings.HasPrefix(ref.Name().Short(), r.GetRefPrefix())
}
