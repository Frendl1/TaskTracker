package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Frendl1/TaskTracker/Logic/Storage"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("  go run cmd/main.go add \"task title\"")
	fmt.Println("  go run cmd/main.go list")
	fmt.Println("  go run cmd/main.go update <id> \"new title\"")
	fmt.Println("  go run cmd/main.go delete <id>")
	fmt.Println("  go run cmd/main.go status <id> <todo|in-progress|done>")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Provide title")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		if err := storage.Add(title); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Added")
		}
	case "list":
		storage.List()
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update <id> \"new title\"")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			return
		}
		title := strings.Join(os.Args[3:], " ")
		if err := storage.UpdateTask(id, title); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Updated")
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			return
		}
		if err := storage.Delete(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Deleted")
		}
	case "status":
		if len(os.Args) < 4 {
			fmt.Println("Usage: status <id> <todo|in-progress|done>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("invalid id")
			return
		}
		status := os.Args[3]
		if err := storage.SetStatus(id, status); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Status updated")
		}
	default:
		usage()
	}
}
