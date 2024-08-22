package assignment2

import (
	"bufio"
	"fmt"
	"os"
)

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
	fmt.Println(firstName, middleName, lastName)

}
