package fixedwidth_test

import (
	"fmt"
	"log"
	"strings"

	fw "github.com/smarty/scanners/v2/fixedwidth"
)

func ExampleScanner() {
	in := strings.Join([]string{
		"name             username",
		"Rob Pike         rob     ",
		"Ken Thompson     ken     ",
		"Robert Griesemer gri     ",
	}, "\n")

	scanner := fw.NewScanner(strings.NewReader(in))

	for scanner.Scan() {
		var (
			name     = scanner.Field(fw.Field(0, 16))
			username = scanner.Field(fw.Field(17, 8))
		)

		fmt.Printf("* % s* %s *\n", name, username)
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	// Output:
	// * name            * username *
	// * Rob Pike        * rob      *
	// * Ken Thompson    * ken      *
	// * Robert Griesemer* gri      *
}

var (
	namef     fw.Substring = func(x string) string { return x[0:16] }
	usernamef fw.Substring = func(x string) string { return x[17:25] }
)

// Define custom [Substring] functions with particular index
// ranges.
func ExampleScanner_substring() {
	in := strings.Join([]string{
		"name             username",
		"Rob Pike         rob     ",
		"Ken Thompson     ken     ",
		"Robert Griesemer gri     ",
	}, "\n")

	scanner := fw.NewScanner(strings.NewReader(in))

	for scanner.Scan() {
		fmt.Printf("* % s* %s *\n", scanner.Field(namef), scanner.Field(usernamef))
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	// Output:
	// * name            * username *
	// * Rob Pike        * rob      *
	// * Ken Thompson    * ken      *
	// * Robert Griesemer* gri      *
}
