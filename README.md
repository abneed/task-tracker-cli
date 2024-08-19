# Task Tracker CLI

A simple command-line application built in Golang to manage tasks. This application allows you to add, update, delete, list, and mark the status of tasks. Tasks are stored in a JSON file for persistence.

# Features
- Add Task: Create a new task with a description.
- Update Task: Modify the description or status of an existing task.
- Delete Task: Remove a task from the list.
- Mark Task Status: Mark a task as in progress or done.
- List Tasks: View all tasks along with their description and status.

# Installation
To install and run the Task Manager CLI, ensure you have Golang installed, then follow these steps:

# Clone the repository
```bash
git clone https://github.com/abneed/task-tracker-cli.git
```

# Navigate to the project directory
```bash
cd task-tracker-cli
```

# Build the application
```bash
go build -o task-cli
```

# Run the application
```bash
./task-cli
```

# Usage
Here’s how you can use the various commands:

# Add a task
```bash
./task-cli add "Your task description"
```

# Update a task
```bash
./task-cli update <task_id> "Updated task description"
```

# Delete a task
```bash
./task-cli delete <task_id>
```

# Marking a task as "in-progress" or "done"
```bash
./task-cli mark-in-progress <task_id>
```

```bash
./task-cli mark-done <task_id>
```

# Listing all tasks
```bash
./task-cli list
```

# Listing all tasks by status 
```bash
./task-cli list todo
```

```bash
./task-cli list in-progress
```

```bash
./task-cli list done
```

# Task Storage
Tasks are stored in a JSON file located in the project directory. This allows for easy persistence and manipulation of task data.

In case that the JSON file doesn't exists, the application will generate that directory and the JSON file storage automatically `db/tasks.json`.


# Contributing
Contributions are welcome! Please fork the repository and submit a pull request for review.
