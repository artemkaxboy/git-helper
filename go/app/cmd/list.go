package cmd

import "fmt"

// ListCmd set of flags and command for list
type ListCmd struct {
	ShortForm bool `long:"short" env:"SHORT" required:"false" description:"short form of the output"`

	CommonOpts
}

// Execute runs list with ListCmd parameters, entry point for "list" command
func (lc *ListCmd) Execute(_ []string) error {
	fmt.Printf("%+v", lc)
	return nil
}
