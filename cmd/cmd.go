package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

type Cmd struct {
	Root     *cobra.Command
	cmds     map[string]*cobra.Command
	duration string
}

func New() *Cmd {
	c := &Cmd{
		Root: &cobra.Command{
			Use: "teamscli",
		},
	}

	c.setup()

	return c
}

func (c *Cmd) Execute() {
	if err := c.Root.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func (c *Cmd) setup() {
	c.cmds = make(map[string]*cobra.Command)
	c.addPresenceCmd()

	for _, ccmd := range c.cmds {
		c.Root.AddCommand(ccmd)
	}
}
