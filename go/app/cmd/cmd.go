package cmd

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
}

// SetCommon satisfies CommonOptionsCommander interface and sets common option fields
// The method called by main for each command
func (co *CommonOpts) SetCommon(_ CommonOpts) {
}
