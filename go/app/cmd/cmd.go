package cmd

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

// CommonOpts sets externally from main, shared across all commands
type CommonOpts struct {
	GitDir    string
	ShortForm bool
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (co *CommonOpts) SetCommon(commonOpts CommonOpts) {
	co.GitDir = commonOpts.GitDir
	co.ShortForm = commonOpts.ShortForm
}
