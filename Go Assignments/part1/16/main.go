package main

import (
	"fmt"
	"log"
	"sync"
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

var printTaskChannel = make(chan string)
var printStatusChannel = make(chan string)
var doneChannel = make(chan bool)
var wg sync.WaitGroup
var globalWg sync.WaitGroup

func printTask() {
	for true {
		log.Println("printTask waiting...")
		key := <-printTaskChannel
		log.Printf("Task[%s]: %v\n", key, Tasks[key])
		printStatusChannel <- key
	}
	log.Println("printTask finished...")
}

func printStatus() {
	for true {
		log.Println("printStatus waiting...")
		key := <-printStatusChannel
		log.Printf("Task status[%s]: %v\n", key, Tasks[key].Status)
		wg.Done()
	}
	log.Println("printStatus finished...")
}

func runList() {
	for key := range Tasks {
		log.Println("runList waiting...")
		fmt.Printf("Key: %s\n", key)
		wg.Add(1)
		printTaskChannel <- key
		wg.Wait()
	}
	log.Println("runList finished...")
	doneChannel <- true
}

func done() {
	log.Println("done waiting...")
	<-doneChannel
	close(printTaskChannel)
	close(printStatusChannel)
	close(doneChannel)
	log.Println("done finished...")
	globalWg.Done()

}

func main() {
	go printTask()
	go printStatus()
	go runList()
	go done()
	globalWg.Add(1)
	globalWg.Wait()

}
