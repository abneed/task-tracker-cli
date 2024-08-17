package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"task-cli/repository"
	"task-cli/service"
)

func main() {
	// Get the command arguments without the command line name
	argumentsWithoutProg := os.Args[1:]

	// Instanciating the Task Service to read/write the "task.json" file
	taskService := service.NewTaskService(
		repository.NewTaskRepository("./db/tasks.json"),
	)

	if len(argumentsWithoutProg) == 0 {
		fmt.Println("Missing arguments")
		return
	}

	switch argument := argumentsWithoutProg[0]; argument {
	case "add":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("description not provided")
		}
		newTaskID, err := taskService.AddTask(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Task added successfully (ID: %d)\n", newTaskID)
	case "update":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("task id not provided")
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal(err)
		}
		description := argumentsWithoutProg[2]
		if len(description) == 0 {
			log.Fatal("description not provided")
		}
		_, err = taskService.UpdateTaskDescription(taskID, description)
		if err != nil {
			log.Fatal(err)
		}
	case "delete":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("task id not provided")
		}

		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal(err)
		}

		taskService.DeleteBy(taskID)
	case "mark-in-progress":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("task id not provided")
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal(err)
		}
		_, err = taskService.UpdateTaskStatus(taskID, "in-progress")
		if err != nil {
			log.Fatal(err)
		}
	case "mark-done":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("task id not provided")
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal(err)
		}
		_, err = taskService.UpdateTaskStatus(taskID, "done")
		if err != nil {
			log.Fatal(err)
		}
	case "list":
		if len(argumentsWithoutProg) > 1 {
			fmt.Printf("%v\n", taskService.GetByStatus(argumentsWithoutProg[1]))
			return
		}
		fmt.Printf("%v\n", taskService.GetAll())
		return
	default:
		fmt.Println("Print help here...")
	}
}
