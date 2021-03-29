package cmd

import (
	"fmt"
	"os"

	"github.com/42wim/teamscli/pkg/presence"
	"github.com/spf13/cobra"
)

func (c *Cmd) addPresenceCmd() {
	c.cmds["PresenceCmd"] = &cobra.Command{
		Use:   "presence",
		Short: "set a specific presence",
		Long:  "possible options are Busy, Available, Offline, DoNotDisturb (DND), Away, BeRightBack (BRB)",
		Run: func(cmd *cobra.Command, args []string) {
			c.Presence(args)
		},
	}
	c.cmds["PresenceCmd"].Flags().StringVarP(&c.duration, "duration", "d", "", "keep presence for specific time (e.g. 1h)")
}

func (c *Cmd) Presence(args []string) {
	if len(args) == 0 {
		c.cmds["PresenceCmd"].Help()
		return
	}

	if c.duration != "" {
		err := presence.SetPresenceDuration(args[0], c.duration)
		if err != nil {
			fmt.Printf("ERROR: setting presence failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("presence set to %s for %s\n", args[0], c.duration)

		return
	}

	err := presence.SetPresence(args[0])
	if err != nil {
		fmt.Printf("ERROR: setting presence failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("presence set to %s\n", args[0])
}
