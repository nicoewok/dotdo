package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nicoewok/dotdo/internal/storage"
	"github.com/nicoewok/dotdo/internal/ui"
	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [task]",
	Short: "Move a task to 'done' status",
	// We use MinimumNArgs(1) so it works WITH or WITHOUT manual quotes
	Args: cobra.MinimumNArgs(1),
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		list, _ := storage.LoadTasks()
		var suggestions []string

		for _, t := range list.Tasks {
			// Suggest anything that isn't already done
			if t.Status != "done" {
				suggestions = append(suggestions, fmt.Sprintf("%d.", t.ID))
				suggestions = append(suggestions, fmt.Sprintf("%d", t.ID))
				// If the title has spaces, we MUST quote it for the shell
				if strings.Contains(t.Title, " ") {
					suggestions = append(suggestions, fmt.Sprintf("\"%s\"", t.Title))
				} else {
					suggestions = append(suggestions, t.Title)
				}
			}
		}
		return suggestions, cobra.ShellCompDirectiveNoFileComp
	},
	Run: func(cmd *cobra.Command, args []string) {
		// Join all args with a space in case the user didn't use quotes manually
		target := strings.Join(args, " ")
		target = strings.Trim(target, "\"")
		target = strings.TrimSpace(target)

		list, _ := storage.LoadTasks()
		found := false
		var matchedTitle string

		for i, t := range list.Tasks {
			if strings.EqualFold(t.Title, target) {
				list.Tasks[i].Status = "done"
				matchedTitle = t.Title
				found = true
				break
			}
		}

		if !found {
			idStr := strings.TrimPrefix(target, "#")
			idStr = strings.TrimSuffix(idStr, ".")
			if id, err := strconv.Atoi(idStr); err == nil && id > 0 {
				for i, t := range list.Tasks {
					if t.ID == id {
						list.Tasks[i].Status = "done"
						matchedTitle = t.Title
						found = true
						break
					}
				}
			}
		}

		if found {
			storage.SaveTasks(list)
			fmt.Printf(" %s Finished: %s\n", ui.GetStatusDot("done"), ui.DoneStyle.Render(matchedTitle))
		} else {
			fmt.Printf(" %s Task not found: %s\n", ui.RedStyle.Render("●"), target)
		}
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
