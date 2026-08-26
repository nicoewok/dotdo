package storage

import (
	"testing"
	"time"
)

func TestNextID(t *testing.T) {
	tests := []struct {
		name     string
		tasks    []Task
		expected int
	}{
		{
			name:     "empty list",
			tasks:    []Task{},
			expected: 1,
		},
		{
			name: "sequential tasks",
			tasks: []Task{
				{ID: 1, Title: "Task 1"},
				{ID: 2, Title: "Task 2"},
			},
			expected: 3,
		},
		{
			name: "non-sequential tasks",
			tasks: []Task{
				{ID: 5, Title: "Task 5"},
				{ID: 2, Title: "Task 2"},
			},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := List{Tasks: tt.tasks}
			got := list.NextID()
			if got != tt.expected {
				t.Errorf("NextID() = %d; want %d", got, tt.expected)
			}
		})
	}
}

func TestResolveDuplicateIDs(t *testing.T) {
	tests := []struct {
		name         string
		initialIDs   []int
		expectedIDs  []int
		wantModified bool
	}{
		{
			name:         "empty list",
			initialIDs:   []int{},
			expectedIDs:  []int{},
			wantModified: false,
		},
		{
			name:         "unique IDs",
			initialIDs:   []int{1, 2, 3},
			expectedIDs:  []int{1, 2, 3},
			wantModified: false,
		},
		{
			name:         "simple duplicate",
			initialIDs:   []int{1, 1, 3},
			expectedIDs:  []int{1, 4, 3},
			wantModified: true,
		},
		{
			name:         "all same ID",
			initialIDs:   []int{2, 2, 2},
			expectedIDs:  []int{2, 3, 4},
			wantModified: true,
		},
		{
			name:         "multiple different duplicates",
			initialIDs:   []int{5, 2, 5, 6, 2},
			expectedIDs:  []int{5, 2, 7, 6, 8},
			wantModified: true,
		},
		{
			name:         "contains zero ID",
			initialIDs:   []int{0, 1, 2},
			expectedIDs:  []int{3, 1, 2},
			wantModified: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tasks []Task
			for _, id := range tt.initialIDs {
				tasks = append(tasks, Task{ID: id, Title: "task"})
			}
			list := List{Tasks: tasks}
			modified := list.ResolveDuplicateIDs()

			if modified != tt.wantModified {
				t.Errorf("ResolveDuplicateIDs() modified = %v; want %v", modified, tt.wantModified)
			}

			if len(list.Tasks) != len(tt.expectedIDs) {
				t.Fatalf("length mismatch: got %d, want %d", len(list.Tasks), len(tt.expectedIDs))
			}

			for i, expectedID := range tt.expectedIDs {
				if list.Tasks[i].ID != expectedID {
					t.Errorf("task[%d].ID = %d; want %d", i, list.Tasks[i].ID, expectedID)
				}
			}

			// Verify all IDs are now unique and positive
			seen := make(map[int]bool)
			for _, task := range list.Tasks {
				if task.ID <= 0 {
					t.Errorf("task has non-positive ID: %d", task.ID)
				}
				if seen[task.ID] {
					t.Errorf("duplicate ID %d still found in result", task.ID)
				}
				seen[task.ID] = true
			}
		})
	}
}

func TestSortByDueDate(t *testing.T) {
	now := time.Now()
	tomorrow := now.Add(24 * time.Hour)
	yesterday := now.Add(-24 * time.Hour)

	list := List{
		Tasks: []Task{
			{ID: 1, Title: "No date 1"},
			{ID: 2, Title: "Tomorrow", Due: tomorrow},
			{ID: 3, Title: "Yesterday", Due: yesterday},
			{ID: 4, Title: "No date 2"},
		},
	}

	list.SortByDueDate()

	expectedOrder := []int{3, 2, 1, 4}
	for i, id := range expectedOrder {
		if list.Tasks[i].ID != id {
			t.Errorf("at index %d: got ID %d, want ID %d", i, list.Tasks[i].ID, id)
		}
	}
}

