package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/nicoewok/dotdo/internal/storage"
)

func setupTestStorage(t *testing.T) (string, func()) {
	t.Helper()
	tempDir, err := os.MkdirTemp("", "dotdo-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}

	// Override HOME so GetStorageDir uses tempDir
	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tempDir)

	dotDir := filepath.Join(tempDir, ".dotdo")
	_ = os.MkdirAll(dotDir, 0755)

	cleanup := func() {
		os.Setenv("HOME", origHome)
		_ = os.RemoveAll(tempDir)
	}

	return tempDir, cleanup
}

func TestDoingAndDoneByIDAndTitle(t *testing.T) {
	_, cleanup := setupTestStorage(t)
	defer cleanup()

	list := storage.List{
		Tasks: []storage.Task{
			{ID: 1, Title: "Buy milk", Status: "todo"},
			{ID: 3, Title: "Feed bunny", Status: "todo"},
		},
	}
	if err := storage.SaveTasksWithoutSync(list); err != nil {
		t.Fatalf("failed to save tasks: %v", err)
	}

	// 1. Move task with ID 3 to doing using "3."
	doingCmd.Run(doingCmd, []string{"3."})
	updatedList, err := storage.LoadTasks()
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	for _, task := range updatedList.Tasks {
		if task.ID == 3 {
			if task.Status != "doing" {
				t.Errorf("task 3 status = %s; want 'doing'", task.Status)
			}
		}
	}

	// 2. Move task with ID 3 to done using bare ID "3"
	doneCmd.Run(doneCmd, []string{"3"})
	updatedList, err = storage.LoadTasks()
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	for _, task := range updatedList.Tasks {
		if task.ID == 3 {
			if task.Status != "done" {
				t.Errorf("task 3 status = %s; want 'done'", task.Status)
			}
		}
	}

	// 3. Move task with title "Buy milk" to doing using "1."
	doingCmd.Run(doingCmd, []string{"1."})
	updatedList, err = storage.LoadTasks()
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	for _, task := range updatedList.Tasks {
		if task.ID == 1 {
			if task.Status != "doing" {
				t.Errorf("task 1 status = %s; want 'doing'", task.Status)
			}
		}
	}

	// 4. Move task with title "Buy milk" to done using title
	doneCmd.Run(doneCmd, []string{"Buy", "milk"})
	updatedList, err = storage.LoadTasks()
	if err != nil {
		t.Fatalf("failed to load tasks: %v", err)
	}

	for _, task := range updatedList.Tasks {
		if task.ID == 1 {
			if task.Status != "done" {
				t.Errorf("task 1 status = %s; want 'done'", task.Status)
			}
		}
	}
}

func TestAddUsesNextID(t *testing.T) {
	_, cleanup := setupTestStorage(t)
	defer cleanup()

	list := storage.List{
		Tasks: []storage.Task{
			{ID: 5, Title: "Existing Task", Status: "todo"},
		},
	}
	_ = storage.SaveTasksWithoutSync(list)

	addCmd.Run(addCmd, []string{"New", "Task"})
	updatedList, _ := storage.LoadTasks()

	if len(updatedList.Tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(updatedList.Tasks))
	}

	if updatedList.Tasks[1].ID != 6 {
		t.Errorf("new task ID = %d; want 6", updatedList.Tasks[1].ID)
	}
}

func TestValidArgsFunctionIncludesNumberSuggestions(t *testing.T) {
	_, cleanup := setupTestStorage(t)
	defer cleanup()

	list := storage.List{
		Tasks: []storage.Task{
			{ID: 1, Title: "Buy milk", Status: "todo"},
			{ID: 2, Title: "Clean room", Status: "doing"},
			{ID: 3, Title: "Already done", Status: "done"},
		},
	}
	_ = storage.SaveTasksWithoutSync(list)

	suggestions, _ := doingCmd.ValidArgsFunction(doingCmd, []string{}, "")
	foundDot1 := false
	foundBare1 := false
	foundTitle1 := false
	for _, s := range suggestions {
		if s == "1." {
			foundDot1 = true
		}
		if s == "1" {
			foundBare1 = true
		}
		if strings.Contains(s, "Buy milk") {
			foundTitle1 = true
		}
	}
	if !foundDot1 {
		t.Errorf("expected suggestions to contain 1., got: %v", suggestions)
	}
	if !foundBare1 {
		t.Errorf("expected suggestions to contain 1, got: %v", suggestions)
	}
	if !foundTitle1 {
		t.Errorf("expected suggestions to contain 'Buy milk', got: %v", suggestions)
	}

	doneSuggestions, _ := doneCmd.ValidArgsFunction(doneCmd, []string{}, "")
	foundDot2 := false
	for _, s := range doneSuggestions {
		if s == "2." {
			foundDot2 = true
		}
	}
	if !foundDot2 {
		t.Errorf("expected done suggestions to contain 2., got: %v", doneSuggestions)
	}
}

func TestRootAndListDoneFlag(t *testing.T) {
	_, cleanup := setupTestStorage(t)
	defer cleanup()

	list := storage.List{
		Tasks: []storage.Task{
			{ID: 1, Title: "Task 1 (todo)", Status: "todo"},
			{ID: 2, Title: "Task 2 (done)", Status: "done"},
		},
	}
	_ = storage.SaveTasksWithoutSync(list)

	// Ensure flags can be toggled
	showDone = false
	rootCmd.Run(rootCmd, []string{})

	showDone = true
	rootCmd.Run(rootCmd, []string{})

	listShowDone = false
	listCmd.Run(listCmd, []string{})

	listShowDone = true
	listCmd.Run(listCmd, []string{})
}

