package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage struct {
	Tasks []Task `json:"tasks"`

}

const fileName = "tasks.json"

func Save(tasks []Task)error{
	data, err:= json.MarshalIndent(tasks,"", "    ")
	if err!=nil{
		fmt.Println( "Ошибка сериализации: ", err)
		return err
	}
	os.WriteFile(fileName, data, 0466)
	return nil
}

func Load() []Task {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return []Task{}
	}
	return tasks
}
