package csv_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/smarty/scanners/csv"
)

func ExampleScanner() {
	in := strings.Join([]string{
		`first_name,last_name,username`,
		`"Rob","Pike",rob`,
		`Ken,Thompson,ken`,
		`"Robert","Griesemer","gri"`,
	}, "\n")
	scanner := csv.NewScanner(strings.NewReader(in))

	for scanner.Scan() {
		fmt.Println(scanner.Record())
	}

	if err := scanner.Error(); err != nil {
		log.Panic(err)
	}

	// Output:
	// [first_name last_name username]
	// [Rob Pike rob]
	// [Ken Thompson ken]
	// [Robert Griesemer gri]
}

// This example shows how csv.Scanner can be configured to handle other
// types of CSV files.
func ExampleScanner_options() {
	in := strings.Join([]string{
		`first_name;last_name;username`,
		`"Rob";"Pike";rob`,
		`# lines beginning with a # character are ignored`,
		`Ken;Thompson;ken`,
		`"Robert";"Griesemer";"gri"`,
	}, "\n")

	scanner := csv.NewScanner(strings.NewReader(in), csv.Comma(';'), csv.Comment('#'))

	for scanner.Scan() {
		fmt.Println(scanner.Record())
	}

	if err := scanner.Error(); err != nil {
		log.Panic(err)
	}

	// Output:
	// [first_name last_name username]
	// [Rob Pike rob]
	// [Ken Thompson ken]
	// [Robert Griesemer gri]
}

// A ColumnScanner maps field values in each row to column
// names.  The column name is taken from the first row, which
// is assumed to be the header row.
func ExampleColumnScanner() {
	in := strings.Join([]string{
		`first_name,last_name,username`,
		`"Rob","Pike",rob`,
		`Ken,Thompson,ken`,
		`"Robert","Griesemer","gri"`,
	}, "\n")
	scanner, _ := csv.NewColumnScanner(strings.NewReader(in))

	for scanner.Scan() {
		fmt.Println(scanner.Column("last_name"), scanner.Column("first_name"))
	}

	if err := scanner.Error(); err != nil {
		log.Panic(err)
	}

	// Output:
	// Pike Rob
	// Thompson Ken
	// Griesemer Robert
}

func ExampleStructScanner() {
	type person struct {
		Firstname string `csv:"first_name"`
		Lastname  string `csv:"last_name"`
		Username  string `csv:"username"`
	}

	in := strings.Join([]string{
		`first_name,last_name,username`,
		`"Rob","Pike",rob`,
		`Ken,Thompson,ken`,
		`"Robert","Griesemer","gri"`,
	}, "\n")

	scanner, _ := csv.NewStructScanner(strings.NewReader(in))

	for scanner.Scan() {
		var p person
		scanner.Populate(&p)
		fmt.Printf("%+v\n", p)
	}

	if err := scanner.Error(); err != nil {
		log.Panic(err)
	}

	// Output:
	// {Firstname:Rob Lastname:Pike Username:rob}
	// {Firstname:Ken Lastname:Thompson Username:ken}
	// {Firstname:Robert Lastname:Griesemer Username:gri}
}
