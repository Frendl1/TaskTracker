package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
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
	//fmt.Println("Задача добавлена!")
}

func (tasks *Tasks) validateIndex(Id int) error{
		if Id < 0 || Id >= len(*tasks){
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
		}
	return nil
}

func (tasks *Tasks) TaskComplate(Id int)error{
	t:= *tasks
	now := time.Now()
	if err:=t.validateIndex(Id); err!=nil{
		return err
	}
	for i:=0; i<len(t); i++{
		if i==Id{
			t[i].Completed = true
			t[i].ComplateAT = &now
		}
	}
	//fmt.Println("Задача выполнена!!")
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
		//fmt.Println("Задача удалена!")
		return  nil
}

func (tasks *Tasks) UpdateTask(Id int, title string){
	t := *tasks
	for i:=0; i< len(t); i++{
		if i==Id{
			t[Id].Title = title
		}

	}
	//fmt.Println("Задача обнавлена")

}

func (tasks *Tasks) TaskList(){
	fmt.Printf("%+v\n\n", tasks)
}

func (tasks *Tasks) CompletedTaskList(){
	t:= *tasks
	var CompletedTask []Task
	for i:=0; i<len(t); i++{
		if t[i].Completed == true{
			CompletedTask = append(CompletedTask, t[i])

		}
	}
		fmt.Printf("%+v\n\n", CompletedTask)

}

func (tasks *Tasks) NotCompletedTaskList(){
	t:= *tasks
	var CompletedTask []Task
	for i:=0; i<len(t); i++{
		if t[i].Completed == false{
			CompletedTask = append(CompletedTask, t[i])

		}
	}
		fmt.Printf("%+v\n\n", CompletedTask)

}


func (tasks *Tasks) Print(){
	table:= table.New(os.Stdout)
	table.AddHeaders("#", "Title", "Status", "Create AT", "Complete AT")
	table.SetRowLines(false)
	for id, t:= range *tasks{
		completed:= "✖️"
		completedAT := ""
		if t.Completed{
			completed = "✓"
			completedAT = t.ComplateAT.Format(time.RFC1123)
		}

		table.AddRow(strconv.Itoa(id), t.Title, completed, t.CreateAT.Format(time.RFC1123), completedAT)
	}
	table.Render()
}