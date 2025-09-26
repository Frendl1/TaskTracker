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
	UpdateAT time.Time `json:"updateAT"`
	CompleteAT *time.Time `json:"completeAT"`
}

/*
func AddTask(desc string)                {}
func UpdateTask(id int, desc string)     {}
func DeleteTask(id int)                  {}
func MarkTask(id int, status string)     {}
func ListTasks()                         {}
func ListTasksByStatus(status string)    {}
*/

func PrimaryId(tasks []Task)int{
	maxID:=0
	for _, t := range tasks{
		if t.Id>maxID{
			maxID = t.Id
		}
	}
	fmt.Println("MaxID:", maxID)
	return maxID+1
}

func add(title string){
	tasks:= Load()
	newID:=PrimaryId(tasks)
	task := Task{
		Title: title,
		Id: newID,
		Status: "todo",
		CreateAT: time.Now(),
		CompleteAT: nil,
	}
	tasks = append(tasks, task)
	fmt.Println("MaxID:", newID)
	fmt.Println("Задача добавлена!")
	err:= Save(tasks)
	if err!= nil{
		fmt.Println("Ошибка сохранения:", err)
		return
	}
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
	now := time.Now()
	if err:=validateIndex(Id); err!=nil{
		return err
	}
		for i:=range t{
		if i==Id{
			t[i].Status = "in-progress" 
			t[i].UpdateAT = now
		}
	}
	
	fmt.Println("Задача в процессе выполнения!!")
	Save(t)
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
	
	fmt.Println("Вы выполнили задачу!!")
	Save(t)
	return nil
}

func delete (id int) error{
	t :=  Load()
	if err:= validateIndex(id); err!= nil{
		return err
	}
	//todos[index] = todos[len(todos)-1]
		//todos = todos[:len(todos)-1]
		t = append(t[:id], t[id+1:]...)
		Save(t)
		fmt.Println("Задача удалена!")
		return  nil
}

func UpdateTask(id int, title string){
	t := Load()
	for i:= range t{
		 //fmt.Println("Цикл начался")
		if i==id{
			//fmt.Println("Айди найдено")
			t[id].Title = title
			t[id].UpdateAT = time.Now()
		}

	}
	fmt.Printf("Задача №%d обнавлена\n", id)
	Save(t)

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


func PrintAll(){
	tasks:= Load()
	table:= table.New(os.Stdout)
	table.AddHeaders("#", "Title", "Status", "Create AT", "UpdateAT", "Complete AT")
	table.SetRowLines(false)
	for id, task:= range tasks{
		completed:= "todo"
		completedAT := ""
		if task.Status=="in-progress"{
			completed = "In progress"
		}
		if task.Status=="done"{
			completed = "Done"
			completedAT = task.CompleteAT.Format(time.RFC1123)
		}
		

		table.AddRow(strconv.Itoa(id), task.Title, completed, task.CreateAT.Format(time.RFC1123),task.UpdateAT.Format(time.RFC1123), completedAT)
	}
	table.Render()
}