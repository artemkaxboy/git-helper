package cmd

// CommonOptionsCommander extends flags.Commander with SetCommon
// All commands should implement this interfaces
type CommonOptionsCommander interface {
	SetCommon(commonOpts CommonOpts)
	Execute(args []string) error
}

// CommonOpts sets externally from main, shared across all commands
type CommonOpts struct {
	GitDir string `long:"git-dir" env:"GITDIR" required:"false" description:"path to git repository"`
	Short  bool   `long:"short" env:"SHORT" required:"false" description:"short form of the output"`
	Filter string `long:"filter" env:"FILTER" required:"false" description:"filter by branch name, last commit author or message"`
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (co *CommonOpts) SetCommon(commonOpts CommonOpts) {
}
