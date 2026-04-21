package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/nicoewok/dotdo/internal/storage"
	"github.com/nicoewok/dotdo/internal/ui"
)

func main() {
	// Print header
	fmt.Print("\033[H\033[2J")
	fmt.Println(ui.GetBunny())
	fmt.Println(lipgloss.NewStyle().Foreground(ui.Grey).Render(" bunny\n  ──────────────"))

	// Load and show tasks
	list, err := storage.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(list.Tasks) == 0 {
		fmt.Println(lipgloss.NewStyle().Foreground(ui.Grey).MarginLeft(2).Render("No tasks for today."))
	} else {
		for _, t := range list.Tasks {
			fmt.Printf("  ○ %s\n", t.Title)
		}
	}
}
