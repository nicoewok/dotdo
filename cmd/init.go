package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nicoewok/dotdo/internal/storage"
	"github.com/nicoewok/dotdo/internal/ui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the dotdo directory and storage",
	Run: func(cmd *cobra.Command, args []string) {
		home, _ := os.UserHomeDir()
		dotFolderPath := filepath.Join(home, ".dotdo")
		configPath := filepath.Join(dotFolderPath, "tasks.json")

		// 1. Create the directory (~/.dotdo)
		if _, err := os.Stat(dotFolderPath); os.IsNotExist(err) {
			err := os.MkdirAll(dotFolderPath, 0755)
			if err != nil {
				fmt.Printf(" %s Failed to create directory: %v\n", ui.RedStyle.Render("●"), err)
				return
			}
			fmt.Printf(" %s Created directory: %s\n", ui.GetStatusDot("todo"), dotFolderPath)
		} else {
			fmt.Printf(" %s Directory already exists: %s\n", ui.GreyStyle.Render("○"), dotFolderPath)
		}

		// 2. Create the tasks.json file if it doesn't exist
		if _, err := os.Stat(configPath); os.IsNotExist(err) {
			emptyList := storage.List{Tasks: []storage.Task{}}
			fileData, _ := json.MarshalIndent(emptyList, "", "  ")

			err := os.WriteFile(configPath, fileData, 0644)
			if err != nil {
				fmt.Printf(" %s Failed to create tasks.json: %v\n", ui.RedStyle.Render("●"), err)
				return
			}
			fmt.Printf(" %s Initialized storage: %s\n", ui.GetStatusDot("todo"), configPath)
		} else {
			fmt.Printf(" %s Storage already exists: %s\n", ui.GreyStyle.Render("○"), configPath)
		}

		// 3. Instruction for Git Sync
		fmt.Println("\n" + ui.GreyStyle.Render(" NEXT STEPS:"))
		fmt.Println(" 1. cd ~/.dotdo")
		fmt.Println(" 2. git init")
		fmt.Println(" 3. git remote add origin <your-private-repo-url>")
		fmt.Println(" 4. dotdo add \"Your first task\"")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
