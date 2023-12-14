package main

import (
	"fmt"
)

// Task struct represents a task
type Task struct {
	ID          int
	Description string
	Status      string
}

// User struct represents a user with tasks
type User struct {
	ID    int
	Name  string
	Tasks []Task
}

// TaskManager struct manages users and tasks
type TaskManager struct {
	Users   []User
	Counter int
}

// AddUser adds a new user to the task manager
func (tm *TaskManager) AddUser(name string) {
	newUser := User{
		ID:   tm.Counter + 1,
		Name: name,
	}

	tm.Users = append(tm.Users, newUser)
	tm.Counter++
	fmt.Println("User added successfully.")
}

// AddTaskToUser adds a new task to a specific user
func (tm *TaskManager) AddTaskToUser(userID int, description string) {
	for i := range tm.Users {
		if tm.Users[i].ID == userID {
			newTask := Task{
				ID:          len(tm.Users[i].Tasks) + 1,
				Description: description,
				Status:      "Pending",
			}
			tm.Users[i].Tasks = append(tm.Users[i].Tasks, newTask)
			fmt.Println("Task added to user successfully.")
			return
		}
	}

	fmt.Println("User not found.")
}

// ViewUserTasks displays all tasks for a specific user
func (tm *TaskManager) ViewUserTasks(userID int) {
	for i := range tm.Users {
		if tm.Users[i].ID == userID {
			fmt.Printf("Tasks for User %d (%s):\n", tm.Users[i].ID, tm.Users[i].Name)
			for _, task := range tm.Users[i].Tasks {
				fmt.Printf("Task ID: %d, Description: %s, Status: %s\n", task.ID, task.Description, task.Status)
			}
			return
		}
	}

	fmt.Println("User not found.")
}

// ... (other methods for updating and deleting tasks, similar to the previous example)

func main() {
	// Create an instance of TaskManager
	taskManager := TaskManager{}

	for {
		fmt.Println("\nTask Manager Menu:")
		fmt.Println("1. Add User")
		fmt.Println("2. Add Task to User")
		fmt.Println("3. View User Tasks")
		// ... (other menu options)
		fmt.Println("6. Exit")
		fmt.Print("Select an option (1-6): ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Print("Enter user name: ")
			fmt.Scanln(&name)
			taskManager.AddUser(name)

		case 2:
			var userID int
			var description string

			fmt.Print("Enter user ID to add task: ")
			fmt.Scanln(&userID)

			fmt.Print("Enter task description: ")
			fmt.Scanln(&description)

			taskManager.AddTaskToUser(userID, description)

		case 3:
			var userID int
			fmt.Print("Enter user ID to view tasks: ")
			fmt.Scanln(&userID)
			taskManager.ViewUserTasks(userID)

		// ... (other cases for additional menu options)

		case 6:
			fmt.Println("Exiting Task Manager. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}
}
