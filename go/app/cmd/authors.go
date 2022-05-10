package cmd

import (
	log "github.com/go-pkgz/lgr"
)

// AuthorsCmd set of flags and command for authors
type AuthorsCmd struct {
	CommonOpts
}

// Execute runs authors with AuthorsCmd parameters, entry point for "authors" command
func (ac *AuthorsCmd) Execute(_ []string) error {

	log.Printf("[DEBUG] execute `authors` on %s", ac.GitDir)

	return nil
}
