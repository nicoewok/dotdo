package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/nicoewok/dotdo/internal/storage"
	"github.com/nicoewok/dotdo/internal/ui"
)

func main() {
	// 1. Clear terminal for that clean Nothing look (optional)
	fmt.Print("\033[H\033[2J")

	// 2. Display the Bunny
	fmt.Println(ui.GetBunny())

	title := lipgloss.NewStyle().
		Foreground(ui.White).
		Bold(true).
		MarginLeft(2).
		Render("D O T D O")
	fmt.Println(title)
	fmt.Println(lipgloss.NewStyle().Foreground(ui.Grey).Render("  ──────────────"))

	// 3. Load and show tasks
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
