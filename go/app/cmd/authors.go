package cmd

import (
	"app/common"
	"app/git"
	"app/output"
	"fmt"
	log "github.com/go-pkgz/lgr"
)

// AuthorsCmd set of flags and command for authors
type AuthorsCmd struct {
	CommonOpts
}

// Execute runs authors with AuthorsCmd parameters, entry point for "authors" command
func (ac *AuthorsCmd) Execute(_ []string) error {

	log.Printf("[DEBUG] execute `authors` on %s", ac.GitDir)

	repo, err := git.Open(ac.GitDir)
	if err != nil {
		return fmt.Errorf("failed to open repository: %w", err)
	}

	remotes, err := repo.GetRemotes()
	if err != nil {
		return fmt.Errorf("failed to read remotes: %w", err)
	}

	for _, remote := range remotes {

		branches, err := remote.GetBranches()
		if err != nil {
			return err
		}

		branchesByAuthor := sumBranchesByAuthor(branches)
		output.PrintAuthors(remote.Name(), branchesByAuthor)
	}

	return nil
}

func sumBranchesByAuthor(branches []*git.RemoteBranch) []common.StringIntPair {
	branchesByAuthor := make(map[string]int)
	for _, branch := range branches {
		branchesByAuthor[branch.LastCommitAuthor("unknown")]++
	}

	pairs := make([]common.StringIntPair, len(branchesByAuthor))
	i := 0
	for k, v := range branchesByAuthor {
		pairs[i] = common.StringIntPair{Key: k, Value: v}
		i++
	}

	return pairs
}
