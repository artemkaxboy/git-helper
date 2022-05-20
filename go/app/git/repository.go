package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"path/filepath"
)

// Repository is a wrapper around go-git repository
type Repository struct {
	Parent   *git.Repository
	rootPath string
}

// Open opens a git repository at the given path
func Open(repositoryPath string) (*Repository, error) {
	absolutePath, err := filepath.Abs(repositoryPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}

	gitRepository, err := git.PlainOpen(absolutePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open repository: %w", err)
	}

	return &Repository{gitRepository, absolutePath}, nil
}
