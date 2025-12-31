package main

import (
	"testing"
)

func TestAddAndList(t *testing.T) {
	// 1. Initialize
	store := MemoryStore{}

	// 2. Action: Add Task
	expectedTitle := "Random Task"
	createdID, err := store.Add(expectedTitle) // Capture the ID!

	if err != nil {
		t.Fatalf("Task could not be added: %v", err)
		// Use Fatalf to stop test immediately if setup fails
	}

	// 3. Action: List Tasks
	taskList, err := store.List() // Use := here

	if err != nil {
		t.Fatalf("Could not list tasks: %v", err)
	}

	// 4. Verification: Check Length (Crucial!)
	if len(taskList) != 1 {
		t.Errorf("Expected 1 task, got %d", len(taskList))
	}

	// 5. Verification: Check Content
	// Since we know length is 1, we can access index [0] directly
	actualTask := taskList[0]

	if actualTask.ID != createdID {
		t.Errorf("Expected ID %d, got %d", createdID, actualTask.ID)
	}

	if actualTask.Title != expectedTitle {
		t.Errorf("Expected Title '%s', got '%s'", expectedTitle, actualTask.Title)
	}
}
