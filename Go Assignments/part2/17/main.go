package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Item struct {
	Id          string
	Title       string
	Description string
}

type Task struct {
	Item   Item
	Status string
}

type Command struct {
	command string
	params  []string
}

var commands = map[string]string{
	"CREATE": "create a task",
	"READ":   "read a task",
	"UPDATE": "update a task",
	"DELETE": "delete a task",
	"EXIT":   "exit",
}

var Tasks map[string]Task
var prompt string

func create(params []string) {

}

func update(params []string) {

}

func read(params []string) {

}

func exit(params []string) {

}

var wg sync.WaitGroup

func waitAndReadCommand() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	run := true
	for run {
		scanner.Scan()
		fullCommand := strings.Trim(scanner.Text(), " ")
		commandSplit := strings.Split(fullCommand, " ")
		if len(commandSplit) < 1 {
			continue
		}
		command := strings.ToUpper(commandSplit[0])
		switch command {
		case "CREATE":
			if len(commandSplit) != 3 {
				outputChannel <- "error: invalid number of parameters [Ex: CREATE <title> <description>]"
				continue
			}
		case "READ":
			if len(commandSplit) != 2 {
				outputChannel <- "error: invalid number of parameters [Ex: READ <id>]"
				continue
			}
		case "UPDATE":
			if len(commandSplit) != 5 {
				outputChannel <- "error: invalid number of parameters [Ex: CREATE <id> <title> <description> <status>]"
				continue
			}
		case "DELETE":
			if len(commandSplit) != 2 {
				outputChannel <- "error: invalid number of parameters [Ex: DELETE <id>]"
				continue
			}
		case "LIST":
			if len(commandSplit) != 1 {
				outputChannel <- "error: invalid number of parameters [Ex: LIST]"
				continue
			}
		case "EXIT":
			run = false
			wg.Add(1)
		default:
			outputChannel <- "error: unknown command"
			continue
		}
		commandChannel <- Command{command: command, params: commandSplit}
		wg.Wait()
	}
}

func processOutput() {
	for {
		output := <-outputChannel
		fmt.Println(output)
		if output == "THE END" {
			wg.Done()
		} else {
			fmt.Print(prompt)
		}
	}
}

func executeCommands() {
	for {
		commandStruct := <-commandChannel

		fmt.Printf("Command: %v\n", commandStruct)

		switch commandStruct.command {
		case "CREATE":
			task := Task{Item: Item{Id: strconv.Itoa(int(time.Now().Unix())), Title: commandStruct.params[1], Description: commandStruct.params[2]}, Status: "false"}
			Tasks[task.Item.Id] = task
			outputChannel <- fmt.Sprintf("Task with Id %s created!", task.Item.Id)
		case "READ":
			task, exists := Tasks[commandStruct.params[1]]
			time.Sleep(10 * time.Second)
			if exists {
				outputChannel <- fmt.Sprintf("Task with Id %s: %v", task.Item.Id, task)
			} else {
				outputChannel <- "read error: unknown task id"
			}

		case "UPDATE":
			_, exists := Tasks[commandStruct.params[1]]
			if exists {
				newTask := Task{Item: Item{Id: commandStruct.params[1], Title: commandStruct.params[2], Description: commandStruct.params[3]}, Status: commandStruct.params[4]}
				Tasks[newTask.Item.Id] = newTask
				outputChannel <- fmt.Sprintf("Task with Id %s updated!", newTask.Item.Id)
			} else {
				outputChannel <- "update error: unknown task id"
			}
		case "DELETE":
			_, exists := Tasks[commandStruct.params[1]]
			if exists {
				delete(Tasks, commandStruct.params[1])
				outputChannel <- fmt.Sprintf("Task with Id %s deleted!", commandStruct.params[1])
			} else {
				outputChannel <- "delete error: unknown task id"
			}
		case "LIST":
			output := fmt.Sprintf("---- LIST ----\n")
			for _, v := range Tasks {
				output += fmt.Sprintf("%v -> %s\n", v.Item, v.Status)
			}
			outputChannel <- output
		case "EXIT":
			outputChannel <- "THE END"
		default:
			outputChannel <- "error: unknown command"
		}
	}
}

var commandChannel = make(chan Command)
var outputChannel = make(chan string)

func main() {
	Tasks = make(map[string]Task)
	prompt = "/>"
	go processOutput()
	go executeCommands()
	waitAndReadCommand()

}
