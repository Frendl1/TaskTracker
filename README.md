Task Tracker CLI

A simple command-line interface (CLI) application for tracking tasks. This application allows users to create, update, and delete tasks, providing a straightforward way to manage personal or team tasks.
Project Structure

task-tracker-cli
├── cmd
│   └── main.go          # Entry point of the application
├── Logic
│   ├── Struct
│   │   └── task.go      # Task struct and methods
│   └── storage
│       └── storage.go   # Task storage interface and implementation
├── go.mod                # Module dependencies
├── go.sum                # Module dependency checksums
└── README.md             # Project documentation

Installation

To install the project, clone the repository and navigate to the project directory:

git clone <repository-url>
cd task_tracker

Then, run the following command to download the dependencies:

go mod tidy

Usage

To run the application, use the following command:

    commands to run .go file

#to show all commands
go run cmd/main.go 
#to add task
go run cmd/main.go add ["Task name"]
#to list all tasks
go run cmd/main.go list 
#to update task by id
go run cmd/main.go update [id] ["New task name"}
#to delete task by id
go run cmd/main.go delete [id]


Features

    Create new tasks
    Update existing tasks
    Delete tasks
    List all tasks

json struct

id: A unique identifier for the task description: A short description of the task status: The status of the task (todo, in-progress, done) createdAt: The date and time when the task was created updatedAt: The date and time when the task was last updated Example: {"id":5,"title":"texttToUpdate","status":true,"createdAt":"2025-09-24T17:38:04.260268+03:00","updatedAt":"2025-09-25T10:01:51.6988301+03:00"}
Contributing

Contributions are welcome! Please open an issue or submit a pull request for any enhancements or bug fixes.
License

This project is licensed under the MIT License. See the LICENSE file for more details.
About

Task tracker is a project used to track and manage your tasks.


