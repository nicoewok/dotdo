package cmd

import (
	"fmt"

	"github.com/nicoewok/dotdo/internal/ui"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Show the dotdo user manual",
	Run: func(cmd *cobra.Command, args []string) {
		// Section Styling
		header := func(s string) string {
			return ui.RedStyle.Bold(true).Render("\n " + s)
		}
		cmdStyle := func(c, desc string) {
			fmt.Printf("  %-25s %s\n", ui.WhiteStyle.Render(c), ui.GreyStyle.Render(desc))
		}

		fmt.Println(header("COMMANDS"))
		cmdStyle("dotdo add [text]", "Add task (no quotes needed)")
		cmdStyle("dotdo add [text] -d [date]", "Add task with YYYY-MM-DD")
		cmdStyle("dotdo doing [tab]", "Focus on a task (by title, 3., or 3)")
		cmdStyle("dotdo done [tab]", "Finish a task (by title, 3., or 3)")
		cmdStyle("dotdo remove", "Purge all 'done' tasks")
		cmdStyle("dotdo sync", "Git pull & push changes")
		cmdStyle("dotdo", "List pending tasks (same as 'dotdo list')")
		cmdStyle("dotdo list", "List pending tasks (-d to include done)")
		cmdStyle("dotdo init", "Initialize storage & folder")

		fmt.Println(header("TIPS"))
		fmt.Println(ui.GreyStyle.Render("  - Running 'dotdo' alone lists pending tasks (same as 'dotdo list')."))
		fmt.Println(ui.GreyStyle.Render("  - Use -d or --done with 'dotdo' or 'dotdo list' to include completed tasks."))
		fmt.Println(ui.GreyStyle.Render("  - Use TAB for autocomplete on titles and numbers."))
		fmt.Println(ui.GreyStyle.Render("  - Expired dates appear in Red."))
		fmt.Println("")
	},
}

func init() {
	// We remove the default help command to use our own
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	rootCmd.AddCommand(helpCmd)
}
