package cli

import (
	"log"

	"github.com/mitchellh/cli"
	"github.com/spiffe/spire/cmd/spire-agent/cli/command"
)

func Run(args []string) int {

	c := cli.NewCLI("spire-agent", "0.0.1") //TODO expose version configuration
	c.Args = args
	c.Commands = map[string]cli.CommandFactory{
		"run": func() (cli.Command, error) {
			return &command.RunCommand{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	return exitStatus
}
