package main

import (
	"errors"
	"fmt"
	"time"
)


type Task struct{
	Title string
	Id int
	Completed bool
	CreateAT time.Time
	ComplateAT *time.Time
}

var id int = 0

type Tasks []Task

func PrimaryId()int{
	id++
	return id
}

func (tasks *Tasks) add(title string){
	task := Task{
		Title: title,
		Id: PrimaryId(),
		Completed: false,
		CreateAT: time.Now(),
		ComplateAT: nil,
	}
	*tasks = append(*tasks, task)
}

func (tasks *Tasks) validateIndex(Id int) error{
		if Id < 0 || Id >= len(*tasks){
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
		}
	return nil
}

func (tasks *Tasks) delete (Id int) error{

	t := *tasks
	
	if err:= t.validateIndex(Id); err!= nil{
		return err
	}

	//todos[index] = todos[len(todos)-1]
		//todos = todos[:len(todos)-1]
		*tasks = append(t[:Id], t[Id+1:]...)
		return  nil
}

func (tasks *Tasks) UpdateTask(Id int, title string){
	t := *tasks
	for i:=0; i< len(t); i++{
		if i==Id{
			t[Id].Title = title
		}

	}

}