package cmd

import (
	"fmt"

	"github.com/nicoewok/dotdo/internal/storage"
	"github.com/nicoewok/dotdo/internal/ui"
	"github.com/spf13/cobra"
)

var listShowDone bool

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tasks",

	// Show bunny logo on every command
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() == "completion" || cmd.Name() == "help" || cmd.Name() == "__complete" {
			return
		}
		if cmd.HasParent() && cmd.Parent().Name() == "completion" {
			return
		}

		fmt.Print("\033[H\033[2J")

		fmt.Println(ui.GetBunny())
	},
	Run: func(cmd *cobra.Command, args []string) {
		list, _ := storage.LoadTasks()

		list.SortByDueDate()

		for _, t := range list.Tasks {
			if t.Status == "done" && !listShowDone {
				continue
			}
			dot := ui.GetStatusDot(t.Status)
			date := ui.FormatDueDate(t.Due)
			if t.Status == "done" {
				fmt.Printf("  %s %s%s\n", dot, ui.DoneStyle.Render(fmt.Sprintf("%d. %s", t.ID, t.Title)), date)
			} else {
				fmt.Printf("  %s %d. %s%s\n", dot, t.ID, t.Title, date)
			}
		}
	},
}

func init() {
	listCmd.Flags().BoolVarP(&listShowDone, "done", "d", false, "Show completed tasks as well")
	rootCmd.AddCommand(listCmd)
}
