package cmd

import (
	"fmt"

	"github.com/nicoewok/dotdo/internal/storage"
	"github.com/nicoewok/dotdo/internal/ui"
	"github.com/spf13/cobra"
)

var showDone bool

var rootCmd = &cobra.Command{
	Use:   "dotdo",
	Short: "Dot-styled todo tool",

	// Show bunny logo on every command
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() == "completion" || cmd.Name() == "help" || cmd.Name() == "__complete" {
			return
		}
		if cmd.HasParent() && cmd.Parent().Name() == "completion" {
			return
		}

		// Ensure storage is initialized by attempting to load tasks (creates storage if needed).
		storage.EnsureInitialized()

		fmt.Print("\033[H\033[2J")

		fmt.Println(ui.GetBunny())
	},
	Run: func(cmd *cobra.Command, args []string) {
		list, _ := storage.LoadTasks()

		list.SortByDueDate()

		var pendingTasks []storage.Task
		var tasksToDisplay []storage.Task
		for _, t := range list.Tasks {
			if t.Status != "done" {
				pendingTasks = append(pendingTasks, t)
				tasksToDisplay = append(tasksToDisplay, t)
			} else if showDone {
				tasksToDisplay = append(tasksToDisplay, t)
			}
		}

		fmt.Printf("%d TASKS PENDING\n", len(pendingTasks))

		for _, t := range tasksToDisplay {
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

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&showDone, "done", "d", false, "Show completed tasks as well")
}
