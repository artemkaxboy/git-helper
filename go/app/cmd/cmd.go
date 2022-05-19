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
	GitDir string
	Filter string
	Age    string
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (co *CommonOpts) SetCommon(opts CommonOpts) {
	if opts.Age == "" {
		opts.Age = "0"
	}
	co.Age = opts.Age
	co.GitDir = opts.GitDir
	co.Filter = opts.Filter
}

func getAge(age string) (*common.LongDuration, error) {
	duration, err := common.ParseLongDuration(age)
	if err != nil {
		return nil, err
	}

	return duration, nil
}
