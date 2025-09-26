package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Frendl1/TaskTracker/Logic/Struct"
	"github.com/aquasecurity/table"
)

const tasksFile = "tasks.json"

func Load() []*Struct.Task {
	var tasks []*Struct.Task

	data, err := os.ReadFile(tasksFile)
	if err != nil {
		// if file missing, return empty slice
		if os.IsNotExist(err) {
			return tasks
		}
		fmt.Println("Ошибка чтения файла:", err)
		return tasks
	}

	if len(data) == 0 {
		return tasks
	}

	if err := json.Unmarshal(data, &tasks); err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		return tasks
	}
	return tasks
}

func Save(tasks []*Struct.Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(tasksFile, data, 0644)
}

func Add(title string) error {
	tasks := Load()
	var maxID int
	for _, t := range tasks {
		if t.Id > maxID {
			maxID = t.Id
		}
	}
	newID := maxID + 1
	now := time.Now()
	task := &Struct.Task{
		Id:        newID,
		Title:     title,
		Status:    "todo",
		CreateAT:  now,
		UpdateAT:  time.Time{},
		CompleteAT: nil,
	}
	tasks = append(tasks, task)
	return Save(tasks)
}

func findIndexByID(tasks []*Struct.Task, id int) int {
	for i, t := range tasks {
		if t.Id == id {
			return i
		}
	}
	return -1
}

func UpdateTask(id int, title string) error {
	tasks := Load()
	idx := findIndexByID(tasks, id)
	if idx == -1 {
		return errors.New("task not found")
	}
	tasks[idx].Title = title
	tasks[idx].UpdateAT = time.Now()
	return Save(tasks)
}

func Delete(id int) error {
	tasks := Load()
	idx := findIndexByID(tasks, id)
	if idx == -1 {
		return errors.New("task not found")
	}
	tasks = append(tasks[:idx], tasks[idx+1:]...)
	return Save(tasks)
}

func SetStatus(id int, status string) error {
	tasks := Load()
	idx := findIndexByID(tasks, id)
	if idx == -1 {
		return errors.New("task not found")
	}
	tasks[idx].Status = status
	now := time.Now()
	tasks[idx].UpdateAT = now
	if status == "done" {
		tasks[idx].CompleteAT = &now
	}
	return Save(tasks)
}

func List() {
    tasks := Load()
    tbl := table.New(os.Stdout) // передаём writer
    tbl.AddHeaders("ID", "Title", "Status", "Created", "Updated", "CompletedAt")
    for _, task := range tasks {
        created := task.CreateAT.Format(time.RFC1123)
        updated := ""
        if !task.UpdateAT.IsZero() {
            updated = task.UpdateAT.Format(time.RFC1123)
        }
        completed := ""
        if task.CompleteAT != nil {
            completed = task.CompleteAT.Format(time.RFC1123)
        }
        tbl.AddRow(strconv.Itoa(task.Id), task.Title, task.Status, created, updated, completed)
    }
    tbl.Render()
}
