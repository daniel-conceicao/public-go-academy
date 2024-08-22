package main

import (
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

func main() {
	var filename string = "thirteen.json"
	var tasks map[string]Task
	file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		log.Fatal(err)
	}
	for _, task := range tasks {
		fmt.Println(task)
	}
}
