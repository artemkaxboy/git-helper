package main

import (
	"app/cmd"
	log "github.com/go-pkgz/lgr"
	"github.com/jessevdk/go-flags"
	"os"
)

// Opts with all cli commands and flags
type Opts struct {
	AuthorsCmd cmd.AuthorsCmd `command:"authors"`
	ListCmd    cmd.ListCmd    `command:"list"`

	GitDir string `long:"git-dir" env:"GIT_DIR" required:"false" description:"The directory where the git repositories are stored"`

	Dbg bool `long:"dbg" env:"DEBUG" description:"debug mode"`
}

func main() {
	var opts Opts
	p := flags.NewParser(&opts, flags.Default)
	setupLog(opts.Dbg)

	p.CommandHandler = func(command flags.Commander, args []string) error {

		// command implements CommonOptionsCommander to allow passing set of extra options defined for all commands
		c := command.(cmd.CommonOptionsCommander)
		c.SetCommon(cmd.CommonOpts{
			GitDir: opts.GitDir,
		})

		err := c.Execute(args)
		if err != nil {
			log.Printf("[ERROR] command failed: %+v", err)
		}
		return err
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}

func setupLog(dbg bool) {
	if dbg {
		log.Setup(log.Debug, log.CallerFile, log.CallerFunc, log.Msec, log.LevelBraces)
		return
	}
	log.Setup(log.Msec, log.LevelBraces)
}
