package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"task-cli/repository"
	"task-cli/service"
	"task-cli/view"
)

func main() {
	// Get the command arguments without the command line name
	argumentsWithoutProg := os.Args[1:]

	// Instanciating the Task Service to handle the tasks actions
	taskService := service.NewTaskService(
		// Instanciating the Task Repository to read/write the "task.json" file
		repository.NewTaskRepository("./db/tasks.json"),
	)

	// Validate if the command was called with arguments, otherwise throw an error and exit
	if len(argumentsWithoutProg) == 0 {
		log.Fatal("task-cli: error: You must provide an option.")
	}

	// Get the first argument and check if is valid
	switch argument := argumentsWithoutProg[0]; argument {
	case "add":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("task-cli: error: description not provided")
		}
		newTaskID, err := taskService.AddTask(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal("task-cli: error: ", err)
		}
		fmt.Printf("Task added successfully (ID: %d)\n", newTaskID)
	case "update":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("task-cli: error: task id not provided")
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal("task-cli: error: ", err)
		}
		description := argumentsWithoutProg[2]
		if len(description) == 0 {
			log.Fatal("task-cli: error: description not provided")
		}
		_, err = taskService.UpdateTaskDescription(taskID, description)
		if err != nil {
			log.Fatal("task-cli: error: ", err)
		}
	case "delete":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("task-cli: error: task id not provided")
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal("task-cli: error: ", err)
		}
		taskService.DeleteBy(taskID)
	case "mark-in-progress":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("ttask-cli: error: ask id not provided")
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal("task-cli: error: ", err)
		}
		_, err = taskService.UpdateTaskStatus(taskID, "in-progress")
		if err != nil {
			log.Fatal("task-cli: error: ", err)
		}
	case "mark-done":
		if len(argumentsWithoutProg) == 1 {
			log.Fatal("task-cli: error: task id not provided")
		}
		taskID, err := strconv.Atoi(argumentsWithoutProg[1])
		if err != nil {
			log.Fatal("task-cli: error: ", err)
		}
		_, err = taskService.UpdateTaskStatus(taskID, "done")
		if err != nil {
			log.Fatal("task-cli: error: ", err)
		}
	case "list":
		if len(argumentsWithoutProg) > 1 {
			view.PromptTableTasks(taskService.GetByStatus(argumentsWithoutProg[1]))
			return
		}
		view.PromptTableTasks(taskService.GetAll())
		return
	default:
		log.Fatal("task-cli: error: option provided not valid")
	}
}
