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
	Id int		 `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"` //todo, in-progress, done 
	CreateAT time.Time `json:"createdAT"`
//	UpdateAT time.Time `json:"updateAT"`
	CompleteAT *time.Time `json:"completeAT"`
}


var id int = 0
type storage Storage
/*
func AddTask(desc string)                {}
func UpdateTask(id int, desc string)     {}
func DeleteTask(id int)                  {}
func MarkTask(id int, status string)     {}
func ListTasks()                         {}
func ListTasksByStatus(status string)    {}
*/

func PrimaryId()int{
	id++
	return id
}

func add(title string){
	tasks:= Load()
	task := Task{
		Title: title,
		Id: PrimaryId(),
		Status: "todo",
		CreateAT: time.Now(),
		CompleteAT: nil,
	}
	tasks = append(tasks, task)
	fmt.Println("Задача добавлена!")
	Save(tasks)
}

func validateIndex(Id int) error{
	tasks := Load()
		if Id < 0 || Id >= len(tasks){
		err := errors.New("Invalid Index")
		fmt.Println(err)
		return err
		}
	return nil
}


func MarkInProgress(Id int)error{
	t:= Load()
	//now := time.Now()
	if err:=validateIndex(Id); err!=nil{
		return err
	}
		for i:=0; i<len(t); i++{
		if i==Id{
			t[i].Status = "in-progress" 
			//t[i].UpdateAT = &now
		}
	}
	
	fmt.Println("Задача в процессе выполнения!!")
	return nil
}

func MarkDone(Id int)error{
	t:= Load()
	now := time.Now()
	if err:=validateIndex(Id); err!=nil{
		return err
	}
		for i:=0; i<len(t); i++{
		if i==Id{
			t[i].Status = "done" 
			t[i].CompleteAT = &now
		}
	}
	
	fmt.Println("Задача в процессе выполнения!!")
	return nil
}

func delete (Id int) error{
	t :=  Load()
	if err:= validateIndex(Id); err!= nil{
		return err
	}
	//todos[index] = todos[len(todos)-1]
		//todos = todos[:len(todos)-1]
		t = append(t[:Id], t[Id+1:]...)
		fmt.Println("Задача удалена!")
		return  nil
}

func UpdateTask(Id int, title string){
	t := Load()
	for i:=0; i< len(t); i++{
		if i==Id{
			t[Id].Title = title
		}

	}
	fmt.Printf("Задача №%d обнавлена\n", Id)

}

func TaskList(){
	print()
}

/*func (tasks *Tasks) CompletedTaskList(){
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
	*/


func Print(){
	t:= Load()
	table:= table.New(os.Stdout)
	table.AddHeaders("#", "Title", "Status", "Create AT", "Complete AT")
	table.SetRowLines(false)
	for id, t:= range t{
		completed:= "todo"
		completedAT := ""
		if t.Status=="in-progress"{
			completed = "In progress"
			completedAT = t.CompleteAT.Format(time.RFC1123)
		}
		if t.Status=="done"{
			completed = "Done"
			completedAT = t.CompleteAT.Format(time.RFC1123)
		}
		

		table.AddRow(strconv.Itoa(id), t.Title, completed, t.CreateAT.Format(time.RFC1123), completedAT)
	}
	table.Render()
}