package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

var Tasks = map[string]Task{
	"1":  {Id: "1", Title: "Title 1", Description: "Description 1", Status: "false"},
	"2":  {Id: "2", Title: "Title 2", Description: "Description 2", Status: "false"},
	"3":  {Id: "3", Title: "Title 3", Description: "Description 3", Status: "false"},
	"4":  {Id: "4", Title: "Title 4", Description: "Description 4", Status: "false"},
	"5":  {Id: "5", Title: "Title 5", Description: "Description 5", Status: "false"},
	"6":  {Id: "6", Title: "Title 6", Description: "Description 6", Status: "false"},
	"7":  {Id: "7", Title: "Title 7", Description: "Description 7", Status: "false"},
	"8":  {Id: "8", Title: "Title 8", Description: "Description 8", Status: "false"},
	"9":  {Id: "9", Title: "Title 9", Description: "Description 9", Status: "false"},
	"10": {Id: "10", Title: "Title 10", Description: "Description 10", Status: "false"},
}

func eleven(tasks ...Task) {

}

func twelve() {
	var filename string = "twelve.json"
	var file *os.File

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, task := range Tasks {
		log.Printf("Task to print:%v", task)
		jsonTask, err := json.Marshal(task)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("JSON to print:%v", jsonTask)
		n4, err := w.Write(jsonTask)
		fmt.Printf("wrote %d bytes\n", n4)
		if err != nil {
			log.Fatal(err)
		}
		w.Flush()
	}

}

func main() {
	for _, task := range Tasks {
		data, err := json.Marshal(task)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", data)
	}

}
