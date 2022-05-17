package cmd

import (
	"app/git"
	"app/output"
	"fmt"
	log "github.com/go-pkgz/lgr"
	"sort"
	"time"
)

// ListCmd set of flags and command for list
type ListCmd struct {
	CommonOpts

	Short bool `short:"s" long:"short" env:"SHORT" required:"false" description:"Short form of the output"`
}

// Execute runs list with ListCmd parameters, entry point for "list" command
func (lc *ListCmd) Execute(_ []string) error {

	log.Printf("[DEBUG] execute `authors` on %s", lc.GitDir)

	repo, err := git.Open(lc.GitDir)
	if err != nil {
		return fmt.Errorf("failed to open repository: %w, make sure that git repository is accessible", err)
	}

	remotes, err := repo.GetRemotes()
	if err != nil {
		return fmt.Errorf("failed to read remotes: %w", err)
	}

	for _, remote := range remotes {

		branches, err := remote.GetBranches(lc.Filter)
		if err != nil {
			return err
		}

		sort.SliceIsSorted(branches, func(i, j int) bool {
			return branches[i].LastCommitTime(time.Now()).After(branches[j].LastCommitTime(time.Now()))
		})

		output.PrintRemoteBranches(remote.Name(), branches, lc.Short)
	}

	return nil
}
