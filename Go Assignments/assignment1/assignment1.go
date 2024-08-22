package main

import (
	"fmt"
	"strings"
)

func main() {
	multipleLineString := `line 1
line 2
line 3`

	multipleLineString = strings.ReplaceAll(multipleLineString, "\n", " ")
	fmt.Print(multipleLineString)

}
