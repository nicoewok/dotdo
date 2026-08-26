package storage

import (
	"encoding/json"
	"os"
	"slices"
	"time"
)

type Task struct {
	ID     int       `json:"id"`
	Title  string    `json:"title"`
	Status string    `json:"status"` // "todo", "doing", "done"
	Due    time.Time `json:"due,omitempty"`
}

type List struct {
	Tasks []Task `json:"tasks"`
}

func (l *List) SortByDueDate() {
	slices.SortFunc(l.Tasks, func(a, b Task) int {
		aHasDue := !a.Due.IsZero()
		bHasDue := !b.Due.IsZero()

		// Case 1: Both have due dates - sort chronologically
		if aHasDue && bHasDue {
			if a.Due.Before(b.Due) {
				return -1
			}
			if a.Due.After(b.Due) {
				return 1
			}
			return 0
		}

		// Case 2: Only 'a' has a due date - 'a' comes first
		if aHasDue && !bHasDue {
			return -1
		}

		// Case 3: Only 'b' has a due date - 'b' comes first
		if !aHasDue && bHasDue {
			return 1
		}

		// Case 4: Neither has a due date - keep relative order
		return 0
	})
}

// NextID returns the next highest available ID for a new task.
func (l *List) NextID() int {
	maxID := 0
	for _, t := range l.Tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

// ResolveDuplicateIDs ensures all tasks have unique, positive IDs.
// If duplicate IDs exist, duplicates are transformed into the next highest number among all tasks.
// Returns true if any task ID was modified.
func (l *List) ResolveDuplicateIDs() bool {
	maxID := 0
	for _, t := range l.Tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}

	seen := make(map[int]bool)
	modified := false

	for i := range l.Tasks {
		if l.Tasks[i].ID <= 0 || seen[l.Tasks[i].ID] {
			maxID++
			l.Tasks[i].ID = maxID
			modified = true
		}
		seen[l.Tasks[i].ID] = true
	}

	return modified
}

// EnsureInitialized checks for the storage folder and file, creating them if missing.
func EnsureInitialized() {
	dir := GetStorageDir()
	path := GetStoragePath()

	// Create directory (~/.dotdo) if missing
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		_ = os.MkdirAll(dir, 0755)
	}

	// Create tasks.json if missing
	if _, err := os.Stat(path); os.IsNotExist(err) {
		empty := List{Tasks: []Task{}}
		data, _ := json.MarshalIndent(empty, "", "  ")
		_ = os.WriteFile(path, data, 0644)
	}
}
