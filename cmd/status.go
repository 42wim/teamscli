package cmd

import (
	"fmt"
	"os"

	"github.com/42wim/teamscli/pkg/status"
	"github.com/spf13/cobra"
)

func (c *Cmd) addStatusCmd() {
	c.cmds["StatusCmd"] = &cobra.Command{
		Use:   "status",
		Short: "set a specific status message",
		Run: func(cmd *cobra.Command, args []string) {
			c.Status(args)
		},
	}
	c.cmds["StatusCmd"].Flags().StringVarP(&c.duration, "duration", "d", "", "keep status for specific time (e.g. 1h)")
	c.cmds["StatusCmd"].Flags().BoolVarP(&c.showonmessage, "show", "s", false, "show status message when people message you")
}

func (c *Cmd) Status(args []string) {
	if len(args) == 0 {
		c.cmds["StatusCmd"].Help()
		return
	}

	if c.duration != "" {
		err := status.SetStatusDuration(args[0], c.duration, c.showonmessage)
		if err != nil {
			fmt.Printf("ERROR: setting status failed: %s\n", err)
			os.Exit(1)
		}

		fmt.Printf("status set to %s for %s\n", args[0], c.duration)

		if c.showonmessage {
			fmt.Println("this message will be shown to people who message you.")
		}

		return
	}

	err := status.SetStatus(args[0], c.showonmessage)
	if err != nil {
		fmt.Printf("ERROR: setting status failed: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("status set to %s\n", args[0])

	if c.showonmessage {
		fmt.Println("this message will be shown to people who message you.")
	}

}
