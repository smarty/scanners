package fields_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/smarty/scanners/fields"
)

// Justification of fields should not affect the scanned values.
func ExampleScanner() {
	in := strings.Join([]string{
		"  a\t  1   foo    i  ",
		"  b\t 10   bar    ii ",
		"  c\t100  bazzle  iii",
	}, "\n")

	scanner := fields.NewScanner(strings.NewReader(in))

	for scanner.Scan() {
		fmt.Println(scanner.Fields())
	}

	if err := scanner.Err(); err != nil {
		log.Panic(err)
	}

	// Output:
	// [a 1 foo i]
	// [b 10 bar ii]
	// [c 100 bazzle iii]
}
