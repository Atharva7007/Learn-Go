package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Completed bool
}

type TaskImporter interface {
	Add(string) (int, error)
	List() ([]Task, error)
	Toggle(int) error
}

type MemoryStore struct {
	tasks []Task
}

func (ms *MemoryStore) Add(title string) (int, error) {
	task := Task{ID: len(ms.tasks), Title: title, Completed: false}
	ms.tasks = append(ms.tasks, task)
	return task.ID, nil
}

func (ms *MemoryStore) List() ([]Task, error) {
	return ms.tasks, nil
}

func (ms *MemoryStore) Toggle(id int) error {
	for _, task := range ms.tasks {
		if task.ID == id {
			task.Completed = !task.Completed
			return nil
		}
	}
	return errors.New("No such task found.")
}

func main() {

	store := MemoryStore{}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Task Manager v1.0 (Type 'quit' to exit)")

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input) // Splits by whitespace

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "add":
			if len(parts) < 2 {
				fmt.Println("Usage: add <title>")
				continue
			}
			title := strings.Join(parts[1:], " ")
			// TODO: Call store.Add(title), handle error, print ID
			id, err := store.Add(title)
			if err != nil {
				fmt.Println("Error adding task:", err)
			} else {
				fmt.Println("Added task with ID:", id)
			}

		case "list":
			// TODO: Call store.List(), iterate and print nicely "[x] ID: Title"
			tasks, err := store.List()
			if err != nil {
				fmt.Println("Error listing tasks:", err)
				continue
			} else {
				for _, task := range tasks {
					fmt.Println("[x] ", task.ID, ": ", task.Title)
				}
			}

		case "toggle":
			if len(parts) < 2 {
				fmt.Println("Usage: toggle <id>")
				continue
			}
			id, _ := strconv.Atoi(parts[1]) // Simple conversion
			// TODO: Call store.Toggle(id), handle error
			err := store.Toggle(id)
			if err != nil {
				fmt.Println("Error toggling task:", err)
			} else {
				fmt.Println("Toggled task with ID:", id)
			}
		case "quit":
			fmt.Println("Bye!")
			return

		default:
			fmt.Println("Unknown command")
		}
	}
}
