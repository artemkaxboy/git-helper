package git

import (
	"github.com/go-git/go-git/v5"
	"path/filepath"
)

type Repository struct {
	Parent   *git.Repository
	rootPath string
}

func Open(repositoryPath string) (*Repository, error) {
	absolutePath, err := filepath.Abs(repositoryPath)
	if err != nil {
		return nil, err
	}

	gitRepository, err := git.PlainOpen(absolutePath)
	if err != nil {
		return nil, err
	}

	return &Repository{gitRepository, absolutePath}, nil
}

func (r *Repository) ToComparableString() string {
	return r.rootPath
}
