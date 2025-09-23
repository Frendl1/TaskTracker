package main

import (
	"fmt"
	"os"
	"strconv"
)

/*
Приложение должно запускаться из командной строки, принимать действия и вводимые пользователем данные в качестве аргументов и сохранять задачи в JSON-файле. Пользователь должен иметь возможность:

    Добавление, обновление и удаление задач

    Отметить задачу как выполняемую или выполненную

    Список всех задач

    Перечислите все выполненные задачи.

    Перечислите все задачи, которые не выполнены

    Перечислите все задачи, которые находятся в процессе выполнения.
*/


func main(){
	args:= os.Args[1:]
	if len(args)< 1{
		fmt.Println("Использование: todo [команда] [аргументы]")
	}
	command := args[0]
	switch command{
		case "add":
			if len(args)<2{
				fmt.Println("Нужно указать описание задачи")
				return
			}
			title:= args[1]
			add(title)
		case "update":
			if len(args)<3{
				fmt.Println("Ошибка: Нужно указать ID и новое описание")
			}
			id, err := strconv.Atoi(args[1])
			if err!=nil{
				fmt.Println("ID должно быть цифрой")
			}
			newDesk:= args[2]
			UpdateTask(id,newDesk)
		case "delete":
			if len(args) < 2{
				fmt.Println("Напишите ID задачи которую хотите удалить")
			}
			id, _ := strconv.Atoi(args[1])
			delete(id)
		case "list":
			TaskList()
		default:
			fmt.Println("Команда не опознана!")
	}


	//todos := Tasks{}
	//storage := NewStorage[Tasks]("tasks.json")
	//storage.Load(&todos)
	/*todos.add("Купить пиво")
	todos.add("НЕ забросить тасктрекер")
	todos.add("ПОпасть в авитоТЕх")
	todos.add("Научиться летать")*/
	//todos.delete(3)
	//todos.delete(1)
	//todos.delete(1)
	//todos.UpdateTask(0,"Бросить ПИть")
	//todos.TaskComplate(2)
	//todos.TaskList()
	//todos.CompletedTaskList()
	//todos.NotCompletedTaskList()
	//todos.Print()
	//storage.Save(todos)


}