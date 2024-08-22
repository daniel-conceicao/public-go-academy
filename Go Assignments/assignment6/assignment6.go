package main

import (
	"assignment6/ageCalculator/age"
	"errors"
	"fmt"
	"log"
	"regexp"

	"gitlab.com/metakeule/fmtdate"
)

func main() {

	//Checks if leap year. Years from 1900 to 9999 are valid. Only dd/MM/yyyy
	re := regexp.MustCompile(`(^(((0[1-9]|1[0-9]|2[0-8])[\/](0[1-9]|1[012]))|((29|30|31)[\\/](0[13578]|1[02]))|((29|30)[\/](0[4,6,9]|11)))[\/](19|[2-9][0-9])\d\d$)|(^29[\/]02[\/](19|[2-9][0-9])(00|04|08|12|16|20|24|28|32|36|40|44|48|52|56|60|64|68|72|76|80|84|88|92|96)$)`)

	fmt.Print("Insert date of birth (DD/MM/YYYY): ")
	var date string
	_, err := fmt.Scanf("%s", &date)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("String entered:", date)
	if !re.MatchString(date) {
		log.Fatal(errors.New("not in the correct format (dd/mm/yyyy)"))
	}
	birthdate, error := fmtdate.Parse("DD/MM/YYYY", date)
	if error != nil {
		log.Fatal(err)
	}
	fmt.Println("Parsed date entered:", birthdate)

	age := age.Age(birthdate)

	fmt.Printf("If you were born on %s you are now %d years old!", fmtdate.Format("DD/MM/YYYY", birthdate), age)
}
