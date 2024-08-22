package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type Pupils struct {
	Pupils []Pupil `json:"pupils"`
}

type Pupil struct {
	FullName    string    `json:"fullName"`
	DateOfBirth time.Time `json:"dateOfBirth"`
	Age         int       `json:"age"`
}

func PrettyPrint(data interface{}) {
	var p []byte
	//    var err := error
	p, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s \n", p)
}

func main() {
	pupilsFile, err := os.Open("pupils.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("no pupils.json file found")
		} else {
			log.Fatal(err)
		}
	}
	defer pupilsFile.Close()

	byteValue, err := io.ReadAll(pupilsFile)
	if err != nil {
		log.Fatal(err)
	}

	var pupilsList Pupils
	err = json.Unmarshal(byteValue, &pupilsList)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of pupils:", len(pupilsList.Pupils))
	//PrettyPrint(pupilsList)
	fmt.Println("Full Name \t\t Date Of Birth \t\t\t Age")
	fmt.Println("------------------------------------------------------------------------------------------")
	for i := 0; i < len(pupilsList.Pupils); i++ {
		fmt.Printf("%s\t\t%v\t\t\t%d\n", pupilsList.Pupils[i].FullName, pupilsList.Pupils[i].DateOfBirth, pupilsList.Pupils[i].Age)
	}

}
