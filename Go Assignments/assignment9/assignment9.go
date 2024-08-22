package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type fullName struct {
	firstName, middleName, lastName string
}

func (fullName fullName) getFullName() string {
	return fmt.Sprintf("%s %s %s", fullName.firstName, fullName.middleName, fullName.lastName)
}

func ReadFromKeyboardWithPrompt(prompt string, scanner *bufio.Scanner) string {
	fmt.Print(prompt)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	firstName := ReadFromKeyboardWithPrompt("First Name: ", scanner)
	middleName := ReadFromKeyboardWithPrompt("Middle Name: ", scanner)
	lastName := ReadFromKeyboardWithPrompt("Last Name: ", scanner)

	fName := fullName{firstName, middleName, lastName}

	fNameSlice := strings.Split(fName.getFullName(), " ")

	fmt.Println("Full name: ", strings.Join(fNameSlice[:], " "))
	fmt.Println("Middle name: ", fNameSlice[1])
	fmt.Println("Surname: ", fNameSlice[2])

}
