package git

import (
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
	"strings"
	"time"
)

type Commit struct {
	Parent *object.Commit
}

func (r *Repository) GetLastCommit(hash plumbing.Hash) (*Commit, error) {
	commit, err := r.Parent.CommitObject(hash)
	if err != nil {
		return nil, err
	}

	return &Commit{Parent: commit}, nil
}

func (c *Commit) Author() string {
	return c.Parent.Author.String()
}

func (c *Commit) Time() time.Time {
	return c.Parent.Author.When
}

func (c *Commit) Title() string {
	title, _, _ := strings.Cut(c.Parent.Message, "\n")
	return title
}
