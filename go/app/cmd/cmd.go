package cmd

import "app/common"

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

// CommonOpts sets externally from main, shared across all commands
type CommonOpts struct {
	GitDir string `short:"d" long:"git-dir" env:"GIT_DIR" required:"false" description:"Path to git repository"`
	Filter string `short:"f" long:"filter" env:"FILTER" required:"false" description:"Filter by branch name, last commit author or message"`
	Age    string `short:"a" long:"age" env:"AGE" required:"false" description:"Minimal age of last commit in branch to show, e.g. 1d, 1w, 1m, 1y" default:"0"`
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (co *CommonOpts) SetCommon(_ CommonOpts) {
}

func getAge(age string) (*common.LongDuration, error) {
	duration, err := common.ParseLongDuration(age)
	if err != nil {
		return nil, err
	}

	return duration, nil
}
